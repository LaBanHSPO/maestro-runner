package emulator

import (
	"os/exec"

	"github.com/devicelab-dev/maestro-runner/pkg/proc"
)

// execCommand is the package-level seam over exec.Command for testability.
// Production code MUST call this variable, not exec.Command directly.
// Tests override via t.Cleanup(func(){ execCommand = exec.Command }) and
// must not run with t.Parallel().
// It defaults to proc.Command so no console window flashes on Windows.
var execCommand = proc.Command

// execLookPath is the companion seam over exec.LookPath. Same rules.
var execLookPath = exec.LookPath
