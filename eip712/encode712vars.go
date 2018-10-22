package eip712

import (
	"math/big"

	"github.com/PaulRBerg/go-ethereum/common"
)

const primaryType = "Mail"

var typesSingle = EIP712Types{
	"EIP712Domain": {
		{
			"name": "name",
			"type": "string",
		},
		{
			"name": "version",
			"type": "string",
		},
		{
			"name": "chainId",
			"type": "uint256",
		},
		{
			"name": "verifyingContract",
			"type": "address",
		},
	},
}
var typesMultiple = EIP712Types{
	"House": {
		{
			"name": "name",
			"type": "string",
		},
	},
	"Person": {
		{
			"name": "name",
			"type": "string",
		},
		{
			"name": "wallet",
			"type": "address",
		},
		{
			"name": "house",
			"type": "House",
		},
	},
	"Mail": {
		{
			"name": "from",
			"type": "Person",
		},
		{
			"name": "to",
			"type": "Person",
		},
		{
			"name": "contents",
			"type": "string",
		},
	},
	"Bail": {
		{
			"name": "from",
			"type": "Person",
		},
		{
			"name": "to",
			"type": "Person",
		},
		{
			"name": "contents",
			"type": "string",
		},
	},
}
var typesStandard = EIP712Types{
	"EIP712Domain": {
		{
			"name": "name",
			"type": "string",
		},
		{
			"name": "version",
			"type": "string",
		},
		{
			"name": "chainId",
			"type": "uint256",
		},
		{
			"name": "verifyingContract",
			"type": "address",
		},
	},
	"Person": {
		{
			"name": "name",
			"type": "string",
		},
		{
			"name": "wallet",
			"type": "address",
		},
	},
	"Mail": {
		{
			"name": "from",
			"type": "Person",
		},
		{
			"name": "to",
			"type": "Person",
		},
		{
			"name": "contents",
			"type": "string",
		},
	},
}
var typesCyclical = EIP712Types{
	"Mail": {
		{
			"name": "from",
			"type": "Mail",
		},
		{
			"name": "to",
			"type": "Mail",
		},
		{
			"name": "contents",
			"type": "string",
		},
	},
}

var domainStandard = EIP712Domain{
	"Ether Mail",
	"1",
	big.NewInt(1),
	common.HexToAddress("0xCcCCccccCCCCcCCCCCCcCcCccCcCCCcCcccccccC"),
	nil,
}

var dataBool = map[string]interface{}{
	"magic": true,
}
var dataStandard = map[string]interface{}{
	"from": map[string]interface{}{
		"name":   "Cow",
		"wallet": "0xCD2a3d9F938E13CD947Ec05AbC7FE734Df8DD826",
	},
	"to": map[string]interface{}{
		"name":   "Bob",
		"wallet": "0xbBbBBBBbbBBBbbbBbbBbbbbBBbBbbbbBbBbbBBbB",
	},
	"contents": "Hello, Bob!",
}
