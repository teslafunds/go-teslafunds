<<<<<<< HEAD
// Copyright 2015 The go-teslafunds Authors
// This file is part of the go-teslafunds library.
//
// The go-teslafunds library is free software: you can redistribute it and/or modify
=======
// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
>>>>>>> 7fdd714... gdbix-update v1.5.0
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
<<<<<<< HEAD
// The go-teslafunds library is distributed in the hope that it will be useful,
=======
// The go-ethereum library is distributed in the hope that it will be useful,
>>>>>>> 7fdd714... gdbix-update v1.5.0
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
<<<<<<< HEAD
// along with the go-teslafunds library. If not, see <http://www.gnu.org/licenses/>.
=======
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.
>>>>>>> 7fdd714... gdbix-update v1.5.0

package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// LoadJSON reads the given file and unmarshals its content.
func LoadJSON(file string, val interface{}) error {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(content, val); err != nil {
		if syntaxerr, ok := err.(*json.SyntaxError); ok {
			line := findLine(content, syntaxerr.Offset)
			return fmt.Errorf("JSON syntax error at %v:%v: %v", file, line, err)
		}
		return fmt.Errorf("JSON unmarshal error in %v: %v", file, err)
	}
	return nil
}

// findLine returns the line number for the given offset into data.
func findLine(data []byte, offset int64) (line int) {
	line = 1
	for i, r := range string(data) {
		if int64(i) >= offset {
			return
		}
		if r == '\n' {
			line++
		}
	}
	return
}
