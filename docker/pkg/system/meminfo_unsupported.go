// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

//go:build !linux && !windows
// +build !linux,!windows

package system // import "github.com/z9905080/dockertest/v3/docker/pkg/system"

// ReadMemInfo is not supported on platforms other than linux and windows.
func ReadMemInfo() (*MemInfo, error) {
	return nil, ErrNotSupportedPlatform
}
