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
	"enode://20897cf56eef845bb2a82dd197c260a90e017ef3941119d6d611df95486e86f1530ca6a61fa0d0bb599ef7f0102ec4cae39661510a92ed5211b2c90393e26678@185.141.62.215:59995",
	"enode://1bc3aa3a284104e6d24a842d0551bb30365ea1d2997cf69d224906a21c41765db048363c864dac3d8e4953516d5ba532cf493e9299dfd77beff23244b08eb53d@185.206.146.57:59995",
	"enode://ad6d7539fe890ae55fefc528f7eede4e56280d3b3a2f3137e6b2b2e09bd00b4cd386bb0f24ec3c31c1a2ae91e4574e53a8a4bc07c5decb4a05cd1e1d645c3e95@185.206.146.59:59995",

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
