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
	discover.MustParseNode("enode://e43c7d4efbe46a870f309b80f88713c10e02f96b883322cfac1134f19c6cff85021c47084e5a534bb16b17bbac34b95fe624fbc050c09f7610653e54eefefe86@185.206.145.206:39993"), //node1
	discover.MustParseNode("enode://d214d972eb0317471e5e7bfcdf42ada30030b02277476443c5dad61b97d4d00fdc79eeec1835e9962251c547d42f3c9342f768e3ae63e2c1782d75bbe9e72e07@185.141.62.215:39993"), //tsf b server
	discover.MustParseNode("enode://a3cb002422cd9496a7c156940ee1b4caccbbd3cd7e34204ddb2cebdb72f5dd91b699082bd823f61bf05bc62b5c0ecfc2fc458fd4d20e5de6bfe8644d3212a275@78.46.67.207:39993"), //tsf europool me

}

// TestNetBootNodes are the enode URLs of the P2P bootstrap nodes running on the
// Morden test network.
var TestNetBootNodes = []*discover.Node{
	// ETH/DEV Go Bootnodes
	/*discover.MustParseNode("enode://n4533109cc9bd7604e4ff6c095f7a1d807e15b38e9bfeb05d3b7c423ba86af0a9e89abbf40bd9dde4250fef114cd09270fa4e224cbeef8b7bf05a51e8260d7uv@198.158.98.185.231:54920"),*/

	// ETH/DEV Cpp Bootnodes
}
