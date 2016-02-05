// Copyright (C) 2014 Andrew Clausen
// This program may be distributed under the BSD-style licence that Go is
// released under; see https://golang.org/LICENSE.
//
// This file implements using an external command to provide a socket-like
// ReadWriter connection.

// +build !windows

package main

import (
	"os/exec"
	"syscall"
)

func setpgidProxy(cmd *exec.Cmd) {
	// Give the proxy its own process group, so it doesn't receive our signals.
	syscall.Setpgid(cmd.Process.Pid, cmd.Process.Pid)
}
