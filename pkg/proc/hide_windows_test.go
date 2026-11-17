//go:build windows

package proc

import (
	"os/exec"
	"syscall"
	"testing"
)

func TestHideSetsNoWindowFlags(t *testing.T) {
	cmd := Command("adb", "devices")
	if cmd.SysProcAttr == nil || !cmd.SysProcAttr.HideWindow || cmd.SysProcAttr.CreationFlags&createNoWindow == 0 {
		t.Fatalf("console window not hidden: %+v", cmd.SysProcAttr)
	}
}

func TestHideKeepsExistingAttrs(t *testing.T) {
	const newProcessGroup = 0x00000200 // CREATE_NEW_PROCESS_GROUP
	cmd := exec.Command("adb")
	cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: newProcessGroup, NoInheritHandles: true}
	Hide(cmd)
	a := cmd.SysProcAttr
	if a.CreationFlags != newProcessGroup|createNoWindow || !a.NoInheritHandles || !a.HideWindow {
		t.Fatalf("existing attrs lost: %+v", a)
	}
}
