// Copyright 2015 The go-ethereum Authors
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

<<<<<<< HEAD
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
=======
// defaultGenesisBlock is a gzip compressed dump of the official default Dubaicoin
>>>>>>> 7fdd714... gdbix-update v1.5.0
// genesis block.
const defaultGenesisBlock = "H4sIAAAAAAAA/62RTU7DMBCF180pIq+7sGMndtiBIsSCS4ztGWopP1VjpKAqd8ckDULQSggxO78375uxfc7yVKwfeofsLmd84l9LCs32a0sMHY4RuuPaVprGNI/lw2Yf4YR9fILxcAXzh9q4OMUTNBDhgt30FxifQxfiKgtT3X9aPhAF99rGt9U035BdmA7/vqcbQm9hvPaIv0hD2w4uRc/LcZHSpbwilK4mJI5GEnlhUZAkR0aVTsuiLgvA+iOXMwstXD6xEjemzftst0tkzxNNKesMWF0Al+BVZWQtVFE7gail0oDW/yDfBC97z9n8DtDgxIlQAgAA"

<<<<<<< HEAD
const ( defaultGenesisBlock = "H4sIAAAAAAAA/62RzWrDMBCEz/FTGJ1zkBz95pZgQg59idVKagT+CbECDsHvXteOSykptKVz08zOp0W6Z/ko0rQNerLNCe3pZ22YIut5JMXadwnq8zwmdKnLg9gv8RkuvklH6E5PMH/QwvV9ukAJCR7YxX+F7iXWMc0203L3EbkYQsRrlW5zqL8y69if/n1RbGNjoXv2ij9oQ1W1OFbv03GyaO8ULQpEENQbq8Ehx6AQhXRWa6NEYDo4ysJ7LycWKnj8Ivv2umGdrVYjGkMhGOWeGWckF9Z6aVigI5mbUHDhpBJ+o9xv0NPqQza8Ac1YNWpUAgAA"

defaultTestnetGenesisBlock = defaultGenesisBlock

)
=======
// defaultTestnetGenesisBlock is a gzip compressed dump of the official default Dubaicoin
// test network genesis block (currently Ropsten).
const defaultTestnetGenesisBlock = "H4sIAAAAAAAA/62RTQrCMBCF9z1FydpFY9okuBMquPAS05ixgSYtbYRK6d2t/RGRCiK+RSDvzXxJJl0QDiKudEqTXUiiNnoVo4JsphJvrG482GoqS2QqmJCHJa6g1s4foclXMD9o4erW15CChxm7+BdoTsYaP9lU8v0zOhtEo66Fv01h/Ia0ps3/fk9VGpdBszbEL7qhKEo1tHbjdrSGRyVsKzEBrhQi0ogh1UBVzLd8WISWPBMI9DGZLiQZFDB/IqcfTutHfB/0dz6EN3P3AQAA"

// defaultDevnetGenesisBlockis a gzip compressed dump of a dev Dubaicoin network genesis block.
const defaultDevnetGenesisBlock = "H4sIAAAAAAAA/62RTQrCMBCF9z1FydpFY9okuBMquPAS05ixgSYtbYRK6d2t/RGRCiK+RSDvzXxJJl0QDiKudEqTXUiiNnoVo4JsphJvrG482GoqS2QqmJCHJa6g1s4foclXMD9o4erW15CChxm7+BdoTsYaP9lU8v0zOhtEo66Fv01h/Ia0ps3/fk9VGpdBszbEL7qhKEo1tHbjdrSGRyVsKzEBrhQi0ogh1UBVzLd8WISWPBMI9DGZLiQZFDB/IqcfTutHfB/0dz6EN3P3AQAA"
>>>>>>> 7fdd714... gdbix-update v1.5.0
