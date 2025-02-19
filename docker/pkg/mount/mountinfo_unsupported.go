// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

//go:build (!windows && !linux && !freebsd) || (freebsd && !cgo)
// +build !windows,!linux,!freebsd freebsd,!cgo

package mount // import "github.com/z9905080/dockertest/v3/docker/pkg/mount"

import (
	"fmt"
	"runtime"
)

func parseMountTable() ([]*Info, error) {
	return nil, fmt.Errorf("mount.parseMountTable is not implemented on %s/%s", runtime.GOOS, runtime.GOARCH)
}
