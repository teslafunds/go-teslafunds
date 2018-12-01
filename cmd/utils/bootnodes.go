// Copyright 2015 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

package utils

import "github.com/teslafunds/go-teslafunds/p2p/discover"

// FrontierBootNodes are the enode URLs of the P2P bootstrap nodes running on
// the Frontier network.
var FrontierBootNodes = []*discover.Node{
	// TSF/DEV Go Bootnodes
	discover.MustParseNode("enode://a9ac2213667c4b80278c7d235d831ab0520ca5de0fa1314c4cc6e51a1e9f936c907ffe72533ed68b54e3ec9423cda32377e9cfec85e97a8169f40e04a33631f0@185.206.145.206:57955"), //node1
	discover.MustParseNode("enode://d214d972eb0317471e5e7bfcdf42ada30030b02277476443c5dad61b97d4d00fdc79eeec1835e9962251c547d42f3c9342f768e3ae63e2c1782d75bbe9e72e07@185.141.62.215:57955"), //tsf b server

}

// TestNetBootNodes are the enode URLs of the P2P bootstrap nodes running on the
// Morden test network.
var TestNetBootNodes = []*discover.Node{
	// ETH/DEV Go Bootnodes
	/*discover.MustParseNode("enode://n4533109cc9bd7604e4ff6c095f7a1d807e15b38e9bfeb05d3b7c423ba86af0a9e89abbf40bd9dde4250fef114cd09270fa4e224cbeef8b7bf05a51e8260d7uv@198.158.98.185.231:54920"),*/

	// ETH/DEV Cpp Bootnodes
}
