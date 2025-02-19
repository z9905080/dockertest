// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

//go:build (!linux && !freebsd) || (freebsd && !cgo)
// +build !linux,!freebsd freebsd,!cgo

package mount // import "github.com/z9905080/dockertest/v3/docker/pkg/mount"

// These flags are unsupported.
const (
	BIND        = 0
	DIRSYNC     = 0
	MANDLOCK    = 0
	NOATIME     = 0
	NODEV       = 0
	NODIRATIME  = 0
	NOEXEC      = 0
	NOSUID      = 0
	UNBINDABLE  = 0
	RUNBINDABLE = 0
	PRIVATE     = 0
	RPRIVATE    = 0
	SHARED      = 0
	RSHARED     = 0
	SLAVE       = 0
	RSLAVE      = 0
	RBIND       = 0
	RELATIME    = 0
	RELATIVE    = 0
	REMOUNT     = 0
	STRICTATIME = 0
	SYNCHRONOUS = 0
	RDONLY      = 0
	mntDetach   = 0
)
