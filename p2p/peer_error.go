<<<<<<< HEAD
// Copyright 2014 The go-ethereum Authors && Copyright 2015 go-teslafunds Authors
// This file is part of the go-teslafunds library.
//
// The go-teslafunds library is free software: you can redistribute it and/or modify
=======
// Copyright 2014 The go-ethereum Authors
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

package p2p

import (
	"fmt"
)

const (
	errInvalidMsgCode = iota
	errInvalidMsg
)

var errorToString = map[int]string{
	errInvalidMsgCode: "invalid message code",
	errInvalidMsg:     "invalid message",
}

type peerError struct {
	code    int
	message string
}

func newPeerError(code int, format string, v ...interface{}) *peerError {
	desc, ok := errorToString[code]
	if !ok {
		panic("invalid error code")
	}
	err := &peerError{code, desc}
	if format != "" {
		err.message += ": " + fmt.Sprintf(format, v...)
	}
	return err
}

func (self *peerError) Error() string {
	return self.message
}

type DiscReason uint

const (
	DiscRequested DiscReason = iota
	DiscNetworkError
	DiscProtocolError
	DiscUselessPeer
	DiscTooManyPeers
	DiscAlreadyConnected
	DiscIncompatibleVersion
	DiscInvalidIdentity
	DiscQuitting
	DiscUnexpectedIdentity
	DiscSelf
	DiscReadTimeout
	DiscSubprotocolError = 0x10
)

var discReasonToString = [...]string{
	DiscRequested:           "Disconnect requested",
	DiscNetworkError:        "Network error",
	DiscProtocolError:       "Breach of protocol",
	DiscUselessPeer:         "Useless peer",
	DiscTooManyPeers:        "Too many peers",
	DiscAlreadyConnected:    "Already connected",
	DiscIncompatibleVersion: "Incompatible P2P protocol version",
	DiscInvalidIdentity:     "Invalid node identity",
	DiscQuitting:            "Client quitting",
	DiscUnexpectedIdentity:  "Unexpected identity",
	DiscSelf:                "Connected to self",
	DiscReadTimeout:         "Read timeout",
	DiscSubprotocolError:    "Subprotocol error",
}

func (d DiscReason) String() string {
	if len(discReasonToString) < int(d) {
		return fmt.Sprintf("Unknown Reason(%d)", d)
	}
	return discReasonToString[d]
}

func (d DiscReason) Error() string {
	return d.String()
}

func discReasonForError(err error) DiscReason {
	if reason, ok := err.(DiscReason); ok {
		return reason
	}
	peerError, ok := err.(*peerError)
	if ok {
		switch peerError.code {
		case errInvalidMsgCode, errInvalidMsg:
			return DiscProtocolError
		default:
			return DiscSubprotocolError
		}
	}
	return DiscSubprotocolError
}
