// Package proc builds subprocess commands that never flash a console window.
//
// A GUI host on Windows (for example a Wails app embedding the runner) has no
// console of its own, so every console child such as adb or emulator gets a
// fresh console window unless it is created with CREATE_NO_WINDOW. All
// subprocesses the runner can start on Windows should go through Command,
// CommandContext or Hide. On other platforms these are plain exec wrappers.
package proc

import (
	"context"
	"os/exec"
)

// Command is exec.Command with the console window hidden.
func Command(name string, args ...string) *exec.Cmd {
	return Hide(exec.Command(name, args...))
}

// CommandContext is exec.CommandContext with the console window hidden.
func CommandContext(ctx context.Context, name string, args ...string) *exec.Cmd {
	return Hide(exec.CommandContext(ctx, name, args...))
}
