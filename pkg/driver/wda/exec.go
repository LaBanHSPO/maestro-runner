package wda

import "github.com/devicelab-dev/maestro-runner/pkg/proc"

// execCommand is the package-level seam over exec.Command so tests can
// fake out external commands like `xcrun simctl`. Production code paths
// added with a test seam MUST call this variable, not exec.Command.
// Tests override via t.Cleanup(func(){ execCommand = exec.Command }) and
// must not run with t.Parallel().
// It defaults to proc.Command so no console window flashes on Windows.
var execCommand = proc.Command
