// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package versions // import "github.com/z9905080/dockertest/v3/docker/types/versions"

import (
	"testing"
)

func assertVersion(t *testing.T, a, b string, result int) {
	if r := compare(a, b); r != result {
		t.Fatalf("Unexpected version comparison result. Found %d, expected %d", r, result)
	}
}

func TestCompareVersion(t *testing.T) {
	assertVersion(t, "1.12", "1.12", 0)
	assertVersion(t, "1.0.0", "1", 0)
	assertVersion(t, "1", "1.0.0", 0)
	assertVersion(t, "1.05.00.0156", "1.0.221.9289", 1)
	assertVersion(t, "1", "1.0.1", -1)
	assertVersion(t, "1.0.1", "1", 1)
	assertVersion(t, "1.0.1", "1.0.2", -1)
	assertVersion(t, "1.0.2", "1.0.3", -1)
	assertVersion(t, "1.0.3", "1.1", -1)
	assertVersion(t, "1.1", "1.1.1", -1)
	assertVersion(t, "1.1.1", "1.1.2", -1)
	assertVersion(t, "1.1.2", "1.2", -1)
}
