//go:build !windows

package proc

import "os/exec"

// Hide is a no-op outside Windows: no console window is ever created there.
func Hide(cmd *exec.Cmd) *exec.Cmd { return cmd }
