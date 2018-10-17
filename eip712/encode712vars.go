package eip712

const primaryType = "Mail"

var typesSingle = map[string]EIP712Type{
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
var typesMultiple = map[string]EIP712Type{
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
var typesStandard = map[string]EIP712Type{
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
var typesCyclical = map[string]EIP712Type{
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
