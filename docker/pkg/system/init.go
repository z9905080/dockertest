// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package system // import "github.com/z9905080/dockertest/v3/docker/pkg/system"

import (
	"syscall"
	"time"
	"unsafe"
)

// Used by chtimes
var maxTime time.Time

func init() {
	// chtimes initialization
	if unsafe.Sizeof(syscall.Timespec{}.Nsec) == 8 {
		// This is a 64 bit timespec
		// os.Chtimes limits time to the following
		maxTime = time.Unix(0, 1<<63-1)
	} else {
		// This is a 32 bit timespec
		maxTime = time.Unix(1<<31-1, 0)
	}
}
