// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

//go:build !windows
// +build !windows

package system // import "github.com/z9905080/dockertest/v3/docker/pkg/system"

// InitLCOW does nothing since LCOW is a windows only feature
func InitLCOW(experimental bool) {
}
