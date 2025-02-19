// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package system // import "github.com/z9905080/dockertest/v3/docker/pkg/system"

import (
	"errors"
)

var (
	// ErrNotSupportedPlatform means the platform is not supported.
	ErrNotSupportedPlatform = errors.New("platform and architecture is not supported")

	// ErrNotSupportedOperatingSystem means the operating system is not supported.
	ErrNotSupportedOperatingSystem = errors.New("operating system is not supported")
)
