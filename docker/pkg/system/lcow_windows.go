// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package system // import "github.com/z9905080/dockertest/v3/docker/pkg/system"

// LCOWSupported returns true if Linux containers on Windows are supported.
func LCOWSupported() bool {
	return lcowSupported
}
