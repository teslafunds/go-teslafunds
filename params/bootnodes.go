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

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// TSF/DEV Go Bootnodes 
	"enode://272cee9bfcc735fe87ad9fae146308a74c6ac4d626d53472571d99fc3c6547f11e25ca6caad58252ed92baf4aaf75943323b22c767480c6fdbad099417a6fbd0@185.141.62.215:59995",
	"enode://f16ef50d66fbd7b907855deb1f8d34c723c7b65674f12c84ce75e102b826797a670cbfb9a9a5a2dbadd23ab4d8bba1131c1d6f5e62614ce499da8092565b8701@185.206.146.57:59995",
	"enode://a3cb002422cd9496a7c156940ee1b4caccbbd3cd7e34204ddb2cebdb72f5dd91b699082bd823f61bf05bc62b5c0ecfc2fc458fd4d20e5de6bfe8644d3212a275@78.46.67.207:59995",
	

}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	
}

// RinkebyV5Bootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network for the experimental RLPx v5 topic-discovery network.
var RinkebyV5Bootnodes = []string{
	
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{

}
