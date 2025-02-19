// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package filters // import "github.com/z9905080/dockertest/v3/docker/types/filters"

func ExampleArgs_MatchKVList() {
	args := NewArgs(
		Arg("label", "image=foo"),
		Arg("label", "state=running"))

	// returns true because there are no values for bogus
	args.MatchKVList("bogus", nil)

	// returns false because there are no sources
	args.MatchKVList("label", nil)

	// returns true because all sources are matched
	args.MatchKVList("label", map[string]string{
		"image": "foo",
		"state": "running",
	})

	// returns false because the values do not match
	args.MatchKVList("label", map[string]string{
		"image": "other",
	})
}
