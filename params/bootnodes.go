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
	// Mainnet Dubaicoin-DBIX Go Bootnodes
	"enode://14a622abe7f93f589b1edca25e52b0a52ac0a2baf1387a197bd44bf5521068bd0c3687d6128afd410eed072b78332ced4ce4a550099b4e76710c9a89ce242088@45.32.65.104:57955?discport=57956", //dbix_node1
	"enode://39bb1c137d4d53017c39b5db3dd1b1dd8e44f710a2b55e79fdc0a1ebd0f139e6973420b1358f7e40f23eb5630b7605d25529fbac05aabdecdc2cfd4d0d2452d8@45.32.220.235:57955?discport=57956", //dbix_node2
	"enode://8cda27e3d924155bc8a823b95e327041161cdc876627921b4a690cae2946580cf0358a8035417f9c0621b4735570e0867171869c1989613a3793c18c65126d18@45.76.35.129:57955", //dbix_node3
	"enode://9073cc28395f32d7a85cf8a0295a353453dd9e682afca7cd968a4e5b5d660385069ecae0f7026916f0f5dda471284114ae6cb8e16d94c69a8fb49733f0b8aa21@45.76.138.199:57955", //dbix_node4
	"enode://1d1affca2ea45105897b2fc2de696cb95dd5dd06516db9a088b9d3ba1deefb81b3a7982ef5b728e09251e005af75c9db33ef1969bee0ed4a6df64b4452039be3@45.32.158.190:57955", //dbix_node5
	"enode://b453c3cf61d945525152b5287de5ad766e91e5456fdc89826b814770737439388077c22c98b1a8a79f0677052d951682af7678e7c87e786d20cd4ad9d1777a82@194.135.95.178:57955", //EU_node1
	"enode://a7717e4a646b21f87a0e25747bbd695893674ffcc65e26e4fa3318690b83620dddf9fb77af4fb30ce9603df325f9a6c1dc5415cf44c3c61ada85ad7e55b07b1c@194.135.95.156:57955", //EU_node2
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Morden test network.
var TestnetBootnodes = []string{
	// Testnet Dubaicoin-DBIX Go Bootnodes
	/*
	"enode://e4533109cc9bd7604e4ff6c095f7a1d807e15b38e9bfeb05d3b7c423ba86af0a9e89abbf40bd9dde4250fef114cd09270fa4e224cbeef8b7bf05a51e8260d6b8@94.242.229.4:57955",
	"enode://8c336ee6f03e99613ad21274f269479bf4413fb294d697ef15ab897598afb931f56beb8e97af530aee20ce2bcba5776f4a312bc168545de4d43736992c814592@94.242.229.203:57955",
	*/
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	/*
	"enode://0cc5f5ffb5d9098c8b8c62325f3797f56509bff942704687b6530992ac706e2cb946b90a34f1f19548cd3c7baccbcaea354531e5983c7d1bc0dee16ce4b6440b@40.118.3.223:30305",
	"enode://1c7a64d76c0334b0418c004af2f67c50e36a3be60b5e4790bdac0439d21603469a85fad36f2473c9a80eb043ae60936df905fa28f1ff614c3e5dc34f15dcd2dc@40.118.3.223:30308",
	"enode://85c85d7143ae8bb96924f2b54f1b3e70d8c4d367af305325d30a61385a432f247d2c75c45c6b4a60335060d072d7f5b35dd1d4c45f76941f62a4f83b6e75daaf@40.118.3.223:30309",
	*/
}
