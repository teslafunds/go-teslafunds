// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package core

// Constants containing the genesis allocation of built-in genesis blocks.
// Their content is an RLP-encoded list of (address, balance) tuples.
// Use mkalloc.go to create/update them.

const mainnetAllocData = "\xf8B\xe0\x94X\xce\xd6$tl\xe5z\u04c42R\u0632\xc1\u03cf\xd3\xfd\u00ca\xd3\xc2\x1b\xce\xcc\xed\xa1\x00\x00\x00\xe0\x94\xbb\xcew\xfe\x16N\xe1\xfb\xb2L\v\xbdc\xf7\u00f3(=\x93\xfb\x8a\xd3\xc2\x1b\xce\xcc\xed\xa1\x00\x00\x00"
const testnetAllocData = "\xf8B\xe0\x94X\xce\xd6$tl\xe5z\u04c42R\u0632\xc1\u03cf\xd3\xfd\u00ca\xd3\xc2\x1b\xce\xcc\xed\xa1\x00\x00\x00\xe0\x94\xbb\xcew\xfe\x16N\xe1\xfb\xb2L\v\xbdc\xf7\u00f3(=\x93\xfb\x8a\xd3\xc2\x1b\xce\xcc\xed\xa1\x00\x00\x00"
const rinkebyAllocData = "\xf8B\xe0\x94X\xce\xd6$tl\xe5z\u04c42R\u0632\xc1\u03cf\xd3\xfd\u00ca\xd3\xc2\x1b\xce\xcc\xed\xa1\x00\x00\x00\xe0\x94\xbb\xcew\xfe\x16N\xe1\xfb\xb2L\v\xbdc\xf7\u00f3(=\x93\xfb\x8a\xd3\xc2\x1b\xce\xcc\xed\xa1\x00\x00\x00"
const devAllocData = "\xf8B\xe0\x94X\xce\xd6$tl\xe5z\u04c42R\u0632\xc1\u03cf\xd3\xfd\u00ca\xd3\xc2\x1b\xce\xcc\xed\xa1\x00\x00\x00\xe0\x94\xbb\xcew\xfe\x16N\xe1\xfb\xb2L\v\xbdc\xf7\u00f3(=\x93\xfb\x8a\xd3\xc2\x1b\xce\xcc\xed\xa1\x00\x00\x00"
