// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

//go:build !linux && !freebsd
// +build !linux,!freebsd

package system // import "github.com/z9905080/dockertest/v3/docker/pkg/system"

import "syscall"

// LUtimesNano is only supported on linux and freebsd.
func LUtimesNano(path string, ts []syscall.Timespec) error {
	return ErrNotSupportedPlatform
}
