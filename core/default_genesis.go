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

const ( defaultGenesisBlock = "H4sIAAAAAAAA/62RzWrDMBCEz/FTGJ1zkBz95pZgQg59idVKagT+CbECDsHvXteOSykptKVz08zOp0W6Z/ko0rQNerLNCe3pZ22YIut5JMXadwnq8zwmdKnLg9gv8RkuvklH6E5PMH/QwvV9ukAJCR7YxX+F7iXWMc0203L3EbkYQsRrlW5zqL8y69if/n1RbGNjoXv2ij9oQ1W1OFbv03GyaO8ULQpEENQbq8Ehx6AQhXRWa6NEYDo4ysJ7LycWKnj8Ivv2umGdrVYjGkMhGOWeGWckF9Z6aVigI5mbUHDhpBJ+o9xv0NPqQza8Ac1YNWpUAgAA"

defaultTestnetGenesisBlock = defaultGenesisBlock

)
