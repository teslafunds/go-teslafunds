<<<<<<< HEAD:tsf/downloader/modes.go
// Copyright 2015 The go-ethereum Authors // go-teslafunds Authors
// This file is part of the go-ethereum library.
//
// The go-teslafunds library is free software: you can redistribute it and/or modify
=======
// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/downloader/modes.go
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
<<<<<<< HEAD:tsf/downloader/modes.go
// The go-teslafunds library is distributed in the hope that it will be useful,
=======
// The go-ethereum library is distributed in the hope that it will be useful,
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/downloader/modes.go
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
<<<<<<< HEAD:tsf/downloader/modes.go
// along with the go-teslafunds library. If not, see <http://www.gnu.org/licenses/>.
=======
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/downloader/modes.go

package downloader

// SyncMode represents the synchronisation mode of the downloader.
type SyncMode int

const (
	FullSync  SyncMode = iota // Synchronise the entire blockchain history from full blocks
	FastSync                  // Quickly download the headers, full sync only at the chain head
	LightSync                 // Download only the headers and terminate afterwards
)
