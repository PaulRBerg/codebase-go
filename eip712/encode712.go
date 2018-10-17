package eip712

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/PaulRBerg/basics/helpers"
	"github.com/PaulRBerg/go-ethereum/common"
	"github.com/PaulRBerg/go-ethereum/common/hexutil"
	"github.com/PaulRBerg/go-ethereum/crypto"
	"math/big"
)

type TypedData struct {
	Types       map[string]EIP712Type `json:"types"`
	PrimaryType string                `json:"primaryType"`
	Domain      EIP712Domain          `json:"domain"`
	Message     EIP712Data            `json:"message"`
}

type EIP712Type []map[string]string

type EIP712TypePriority struct {
	Type  string
	Value uint
}

type EIP712Data = map[string]interface{}

type EIP712Domain struct {
	Name              string         `json:"name"`
	Version           string         `json:"version"`
	ChainId           *big.Int       `json:"chainId"`
	VerifyingContract common.Address `json:"verifyingContract"`
	Salt              hexutil.Bytes  `json:"salt"`
}

func MainEncode712() {
	types := typesStandard
	data := dataStandard

	if primaryType != "" && types[primaryType] == nil {
		panic(fmt.Errorf("primaryType specified but undefined"))
	}

	hash := hashStruct(types, "primaryType", data, 0)
	fmt.Printf("hash: %s\n", hash.String())
}

// `encode(domainSeparator : 𝔹²⁵⁶, message : 𝕊) = "\x19\x01" ‖ domainSeparator ‖ hashStruct(message)`
func hashStruct(types map[string]EIP712Type, key string, data EIP712Data, depth int) common.Hash {
	helpers.PrintJson("hashStruct", map[string]interface{}{
		"depth": depth,
	})
	//fmt.Printf("hashStruct: depth %d\n\n", depth)

	typeEncoding := encodeType(types)
	typeHash := hex.EncodeToString(crypto.Keccak256([]byte(typeEncoding)))

	dataEncoding := encodeData(types, key, data, depth)
	dataHash := hex.EncodeToString(crypto.Keccak256([]byte(dataEncoding)))

	var buffer bytes.Buffer
	buffer.WriteString(typeHash)
	buffer.WriteString(dataHash)
	hash := common.BytesToHash(crypto.Keccak256(buffer.Bytes()))

	if depth == 0 {
		fmt.Printf("typeEncoding %s\n", typeEncoding)
		fmt.Printf("dataEncoding %s\n", dataEncoding)
	}
	return hash
}
