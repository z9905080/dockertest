// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

//go:build !windows
// +build !windows

package system // import "github.com/z9905080/dockertest/v3/docker/pkg/system"

import (
	"time"
)

// setCTime will set the create time on a file. On Unix, the create
// time is updated as a side effect of setting the modified time, so
// no action is required.
func setCTime(path string, ctime time.Time) error {
	return nil
}
