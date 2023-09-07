// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

//go:build !windows
// +build !windows

package archive // import "github.com/z9905080/dockertest/v3/docker/pkg/archive"

import (
	"path/filepath"
)

func normalizePath(path string) string {
	return filepath.ToSlash(path)
}
