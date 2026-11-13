package cli

import (
	"context"
	"os"
	"testing"

	"github.com/devicelab-dev/maestro-runner/pkg/executor"
	"github.com/devicelab-dev/maestro-runner/pkg/flow"
	"github.com/devicelab-dev/maestro-runner/pkg/report"
)

func parseMockFlow(t *testing.T, dir, body string) flow.Flow {
	t.Helper()
	path := dir + "/flow.yaml"
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	f, err := flow.ParseFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return *f
}

func TestExecuteFlowsWithDriverContext_OptsReplaceCallbacks(t *testing.T) {
	dir := t.TempDir()
	cfg := &RunConfig{Platform: "mock", Devices: []string{"mock-device"}, OutputDir: dir + "/reports"}
	driver, cleanup, err := CreateDriver(cfg)
	if err != nil {
		t.Fatalf("CreateDriver: %v", err)
	}
	defer cleanup()
	f := parseMockFlow(t, dir, "- tapOn: \"A\"\n- tapOn: \"B\"\n")

	var steps int
	result, err := ExecuteFlowsWithDriverContext(context.Background(), driver, cfg, []flow.Flow{f}, func(rc *executor.RunnerConfig) {
		rc.OnFlowStart = nil
		rc.OnFlowEnd = nil
		rc.OnStepComplete = func(int, string, bool, int64, string) { steps++ }
	})
	if err != nil {
		t.Fatalf("run: %v", err)
	}
	if result.Status != report.StatusPassed {
		t.Fatalf("status = %v, want passed", result.Status)
	}
	if steps != 2 {
		t.Errorf("custom OnStepComplete called %d times, want 2", steps)
	}
}

func TestExecuteFlowsWithDriverContext_CancelledContextStopsRun(t *testing.T) {
	dir := t.TempDir()
	cfg := &RunConfig{Platform: "mock", Devices: []string{"mock-device"}, OutputDir: dir + "/reports"}
	driver, cleanup, err := CreateDriver(cfg)
	if err != nil {
		t.Fatalf("CreateDriver: %v", err)
	}
	defer cleanup()
	f := parseMockFlow(t, dir, "- tapOn: \"A\"\n")

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	var steps int
	result, err := ExecuteFlowsWithDriverContext(ctx, driver, cfg, []flow.Flow{f}, func(rc *executor.RunnerConfig) {
		rc.OnFlowStart = nil
		rc.OnFlowEnd = nil
		rc.OnStepComplete = func(_ int, _ string, passed bool, _ int64, _ string) {
			if passed {
				steps++
			}
		}
	})
	if err == nil && result != nil && result.Status == report.StatusPassed {
		t.Fatal("a cancelled context must not produce a passed run")
	}
	if steps != 0 {
		t.Errorf("%d steps passed after cancel, want 0", steps)
	}
}
