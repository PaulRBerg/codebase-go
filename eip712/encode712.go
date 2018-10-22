package eip712

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/PaulRBerg/basics/helpers"
	"github.com/PaulRBerg/go-ethereum/common"
	"github.com/PaulRBerg/go-ethereum/common/hexutil"
	"github.com/PaulRBerg/go-ethereum/crypto"
)

type TypedData struct {
	Types       EIP712Types  `json:"types"`
	PrimaryType string       `json:"primaryType"`
	Domain      EIP712Domain `json:"domain"`
	Message     EIP712Data   `json:"message"`
}

type EIP712Type []map[string]string
type EIP712Types map[string]EIP712Type

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

const (
	TypeArray   = "array"
	TypeAddress = "address"
	TypeBool    = "bool"
	TypeBytes   = "bytes"
	TypeInt     = "int"
	TypeMap     = "map"
	TypeString  = "string"
)

var typedData = TypedData{
	typesStandard,
	primaryType,
	domainStandard,
	dataStandard,
}

func MainEncode712() {
	typeEncoding := typedData.encodeTypeV2(typedData.PrimaryType)
	fmt.Printf("encodeType: %s\n", string(typeEncoding)) // should be `Mail(Person from,Person to,string contents)Person(string name,address wallet)`

	mailTypeHash := typedData.typeHash("Mail")
	fmt.Printf("mailTypeHash: 0x%s\n", common.Bytes2Hex(mailTypeHash)) // should be `0xa0cedeb2dc280ba39b857546d74f5549c3a1d7bdc2dd96bf881f76108e23dac2`

	dataEncoding := typedData.encodeDataV2(typedData.PrimaryType, typedData.Message)
	fmt.Printf("encodeData: 0x%s\n", common.Bytes2Hex(dataEncoding)) // should be `0xa0cedeb2dc280ba39b857546d74f5549c3a1d7bdc2dd96bf881f76108e23dac2fc71e5fa27ff56c350aa531bc129ebdf613b772b6604664f5d8dbe21b85eb0c8cd54f074a4af31b4411ff6a60c9719dbd559c221c8ac3492d9d872b041d703d1b5aadf3154a261abdd9086fc627b61efca26ae5702701d05cd2305f7c52a2fc8`

	signature := Sign()
	fmt.Printf("signature: 0x%s\n", common.Bytes2Hex(signature))
}

func Sign() []byte {
	buffer := bytes.Buffer{}
	buffer.WriteString("\x19")
	buffer.WriteString("\x01")
	buffer.Write(hashStruct("EIP712Domain", typedData.Domain.Map()))
	buffer.Write(hashStruct(typedData.PrimaryType, typedData.Message))
	signature := crypto.Keccak256(buffer.Bytes())
	return signature
}

// hashStruct generates the following encoding for the given domain and message:
// `encode(domainSeparator : 𝔹²⁵⁶, message : 𝕊) = "\x19\x01" ‖ domainSeparator ‖ hashStruct(message)`
func hashStruct(primaryType  string, data EIP712Data) []byte {
	helpers.PrintJson("hashStruct", map[string]interface{}{
		"primaryType": primaryType,
	})

	return crypto.Keccak256(typedData.encodeDataV2(primaryType, data))
}

/*
function structHash(primaryType, data) {
	return ethUtil.keccak256(encodeData(primaryType, data));
}
 */

// Map is a helper function to generate a map version of the domain
func (domain *EIP712Domain) Map() map[string]interface{} {
	dataMap := map[string]interface{}{
		"chainId": domain.ChainId,
	}

	if len(domain.Name) > 0 {
		dataMap["name"] = domain.Name
	}

	if len(domain.Version) > 0 {
		dataMap["version"] = domain.Version
	}

	if len(domain.VerifyingContract) > 0 {
		dataMap["verifyingContract"] = domain.VerifyingContract
	}

	if len(domain.Salt) > 0 {
		dataMap["salt"] = domain.Salt
	}
	return dataMap
}