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

package errs

import (
	"fmt"

<<<<<<< HEAD
	"github.com/teslafunds/go-teslafunds/logger"
	"github.com/teslafunds/go-teslafunds/logger/glog"
=======
	"github.com/dubaicoin-dbix/go-dubaicoin/logger/glog"
>>>>>>> 7fdd714... gdbix-update v1.5.0
)

/*
Errors implements an error handler providing standardised errors for a package.
Fields:

 Errors:
  a map from error codes to description

 Package:
  name of the package/component
*/
type Errors struct {
	Errors  map[int]string
	Package string
}

/*
Error implements the standard go error interface.

  errors.New(code, format, params ...interface{})

Prints as:

 [package] description: details

where details is fmt.Sprintf(self.format, self.params...)
*/
type Error struct {
	Code    int
	Name    string
	Package string
	message string
	format  string
	params  []interface{}
}

func (self *Errors) New(code int, format string, params ...interface{}) *Error {
	name, ok := self.Errors[code]
	if !ok {
		panic("invalid error code")
	}
	return &Error{
		Code:    code,
		Name:    name,
		Package: self.Package,
		format:  format,
		params:  params,
	}
}

func (self Error) Error() (message string) {
	if len(message) == 0 {
		self.message = fmt.Sprintf("[%s] ERROR: %s", self.Package, self.Name)
		if self.format != "" {
			self.message += ": " + fmt.Sprintf(self.format, self.params...)
		}
	}
	return self.message
}

func (self Error) Log(v glog.Verbose) {
	if v {
		v.Infoln(self)
	}
}
