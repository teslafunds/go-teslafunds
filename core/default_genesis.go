// Copyright 2014 The go-ethereum Authors
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

import (
    "compress/gzip"
    "encoding/base64"
    "io"
    "strings"
)

func NewDefaultGenesisReader() (io.Reader, error) {
    return gzip.NewReader(base64.NewDecoder(base64.StdEncoding, strings.NewReader(defaultGenesisBlock)))
}

// defaultGenesisBlock is a gzip compressed dump of the official default Teslafunds
// genesis block.

const ( defaultGenesisBlock = "H4sIAAAAAAAA/61Ry2rDMBA8x19hdM5BsrSWmluLCT3kJ1YruRH4EWIFXIL/PY4dh1BSaEvnppmd2UF7TtIRrGkb8myTMt7zR0ih2XoeiaH2XcT6MI+BKUyxhbdFPuDRN/Edu/2TmD9gyfV9PGKBEW+xC/+B3S7UIc60MPnrXXKhLAOdqvg5i+ZLZB36/b/3pDY0Frtnn/gDN1ZVS6P1PD0nivekx+YKlZAlmJwkeMrRKlFmYJ20qKzIufBw9aXMYoW3I8J324Z1slpdj6cBMiU1GZs5lC/IjQBwXmifGSXRGYsklftF8lR8SIYLarfa9lECAAA="

defaultTestnetGenesisBlock = defaultGenesisBlock

)
