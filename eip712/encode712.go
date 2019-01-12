package eip712

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
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

var typedData = TypedData{
	typesStandard,
	primaryType,
	domainStandard,
	dataStandard,
}

func MainEncode712() {
	// encodeType
	mailTypeEncoding := string(typedData.encodeType(typedData.PrimaryType))
	if mailTypeEncoding != "Mail(Person from,Person to,string contents)Person(string name,address wallet)" {
		panic(fmt.Errorf("mailTypeEncoding %s is incorrect", mailTypeEncoding))
	}
	fmt.Printf("mailTypeEncoding: %s\n", mailTypeEncoding) // should be `Mail(Person from,Person to,string contents)Person(string name,address wallet)`

	// encodeType
	domainTypeEncoding := string(typedData.encodeType("EIP712Domain"))
	if domainTypeEncoding != "EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)" {
		panic(fmt.Errorf("domainTypeEncoding %s is incorrect", domainTypeEncoding))
	}
	fmt.Printf("domainTypeEncoding: %s\n", domainTypeEncoding)

	// typeHash
	mailTypeHash := fmt.Sprintf("0x%s", common.Bytes2Hex(typedData.typeHash(typedData.PrimaryType)))
	if mailTypeHash != "0xa0cedeb2dc280ba39b857546d74f5549c3a1d7bdc2dd96bf881f76108e23dac2" {
		panic(fmt.Errorf("mailTypeHash %s is incorrect", mailTypeHash))
	}
	fmt.Printf("mailTypeHash: %s\n", mailTypeHash)

	// encodeData
	dataEncoding := fmt.Sprintf("0x%s", common.Bytes2Hex(typedData.encodeData(typedData.PrimaryType, typedData.Message)))
	if dataEncoding != "0xa0cedeb2dc280ba39b857546d74f5549c3a1d7bdc2dd96bf881f76108e23dac2fc71e5fa27ff56c350aa531bc129ebdf613b772b6604664f5d8dbe21b85eb0c8cd54f074a4af31b4411ff6a60c9719dbd559c221c8ac3492d9d872b041d703d1b5aadf3154a261abdd9086fc627b61efca26ae5702701d05cd2305f7c52a2fc8" {
		panic(fmt.Errorf("dataEncoding %s is incorrect", dataEncoding))
	}
	fmt.Printf("dataEncoding: %s\n", dataEncoding)

	// hashStruct
	mainHash := fmt.Sprintf("0x%s", common.Bytes2Hex(hashStruct(typedData.PrimaryType, typedData.Message)))
	if mainHash != "0xc52c0ee5d84264471806290a3f2c4cecfc5490626bf912d01f240d7a274b371e" {
		panic(fmt.Errorf("mainHash %s is incorrect", mainHash))
	}
	fmt.Printf("mainHash: %s\n", mainHash)

	// hashStruct
	domainHash := fmt.Sprintf("0x%s", common.Bytes2Hex(hashStruct("EIP712Domain", typedData.Domain.Map())))
	if domainHash != "0xf2cee375fa42b42143804025fc449deafd50cc031ca257e0b194a650a912090f" {
		panic(fmt.Errorf("domainHash %s is incorrect", domainHash))
	}
	fmt.Printf("domainHash: %s\n", domainHash)

	// signature
	sigHash := fmt.Sprintf("0x%s", common.Bytes2Hex(signatureHash()))
	if sigHash != "0xbe609aee343fb3c4b28e1df9e632fca64fcfaede20f02e86244efddf30957bd2" {
		panic(fmt.Errorf("sigHash %s is incorrect", sigHash))
	}
	fmt.Printf("sigHash: %s\n", sigHash)
}

func signatureHash() []byte {
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
func hashStruct(primaryType string, data EIP712Data) []byte {
	return crypto.Keccak256(typedData.encodeData(primaryType, data))
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
