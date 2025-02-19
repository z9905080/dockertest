// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package homedir // import "github.com/z9905080/dockertest/v3/docker/pkg/homedir"

import (
	"os"

	"github.com/z9905080/dockertest/v3/docker/pkg/idtools"
)

// GetStatic returns the home directory for the current user without calling
// os/user.Current(). This is useful for static-linked binary on glibc-based
// system, because a call to os/user.Current() in a static binary leads to
// segfault due to a glibc issue that won't be fixed in a short term.
// (#29344, golang/go#13470, https://sourceware.org/bugzilla/show_bug.cgi?id=19341)
func GetStatic() (string, error) {
	uid := os.Getuid()
	usr, err := idtools.LookupUID(uid)
	if err != nil {
		return "", err
	}
	return usr.Home, nil
}
