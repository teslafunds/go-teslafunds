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

package node

import (
	"fmt"
	"reflect"
)

// DuplicateServiceError is returned during Node startup if a registered service
// constructor returns a service of the same type that was already started.
type DuplicateServiceError struct {
	Kind reflect.Type
}

// Error generates a textual representation of the duplicate service error.
func (e *DuplicateServiceError) Error() string {
	return fmt.Sprintf("duplicate service: %v", e.Kind)
}

// StopError is returned if a Node fails to stop either any of its registered
// services or itself.
type StopError struct {
	Server   error
	Services map[reflect.Type]error
}

// Error generates a textual representation of the stop error.
func (e *StopError) Error() string {
	return fmt.Sprintf("server: %v, services: %v", e.Server, e.Services)
}
