// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package archive // import "github.com/z9905080/dockertest/v3/docker/pkg/archive"

// Whiteouts are files with a special meaning for the layered filesystem.
// Docker uses AUFS whiteout files inside exported archives. In other
// filesystems these files are generated/handled on tar creation/extraction.

// WhiteoutPrefix prefix means file is a whiteout. If this is followed by a
// filename this means that file has been removed from the base layer.
const WhiteoutPrefix = ".wh."

// WhiteoutMetaPrefix prefix means whiteout has a special meaning and is not
// for removing an actual file. Normally these files are excluded from exported
// archives.
const WhiteoutMetaPrefix = WhiteoutPrefix + WhiteoutPrefix

// WhiteoutLinkDir is a directory AUFS uses for storing hardlink links to other
// layers. Normally these should not go into exported archives and all changed
// hardlinks should be copied to the top layer.
const WhiteoutLinkDir = WhiteoutMetaPrefix + "plnk"

// WhiteoutOpaqueDir file means directory has been made opaque - meaning
// readdir calls to this directory do not follow to lower layers.
const WhiteoutOpaqueDir = WhiteoutMetaPrefix + ".opq"
