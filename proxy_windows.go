// Copyright (C) 2014 Andrew Clausen
// This program may be distributed under the BSD-style licence that Go is
// released under; see https://golang.org/LICENSE.
//
// This file implements using an external command to provide a socket-like
// ReadWriter connection.

// +build windows

package main

func setpgidProxy() {
	// FIXME Windows doesn't have this syscall - ignore it for now
}