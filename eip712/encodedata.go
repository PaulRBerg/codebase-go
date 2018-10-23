package eip712

import (
	"bytes"
	"github.com/PaulRBerg/go-ethereum/accounts/abi"
	"math/big"
	"reflect"
	"strconv"
	"strings"

	"github.com/PaulRBerg/go-ethereum/common"
	"github.com/PaulRBerg/go-ethereum/crypto"
)

func (typedData *TypedData) typeHash(primaryType string) []byte {
	return crypto.Keccak256(typedData.encodeType(primaryType))
}

// encodeData generates the following encoding:
// `enc(value₁) ‖ enc(value₂) ‖ … ‖ enc(valueₙ)`
//
// each encoded member is 32-byte long
func (typedData *TypedData) encodeData(primaryType string, data map[string]interface{}) []byte {
	encTypes := []string{}
	encValues := []interface{}{}

	// Add typehash
	encTypes = append(encTypes, "bytes32")
	encValues = append(encValues, typedData.typeHash(primaryType))

	// Handle primitive values
	handlePrimitiveValue := func(_type string, value interface{}, primaryType string, data interface{}) (string, interface{}) {
		var encType string
		var encValue interface{}

		if _type == "address" {
			encType = "address"
			bytesValue := []byte{}
			for i := 0; i < 12; i++ {
				bytesValue = append(bytesValue, 0)
			}
			for _, _byte := range value.(common.Address).Bytes() {
				bytesValue = append(bytesValue, _byte)
			}
			encValue = bytesValue
		} else if _type == "bytes" {
			encType = "bytes32"
			value := crypto.Keccak256(value.([]byte))
			encValue = value
		} else if _type == "string" {
			encType = "bytes32"
			value := crypto.Keccak256([]byte(value.(string)))
			encValue = value
		} else if _type == "bool" {
			encType = "uint256"
			var int64Val int64
			if value.(bool) {
				int64Val = 1
			}
			encValue = abi.U256(big.NewInt(int64Val))
		} else if strings.HasPrefix(_type, "bytes") {
			encTypes = append(encTypes, "bytes32")
			sizeStr := strings.TrimPrefix(_type, "bytes")
			size, _ := strconv.Atoi(sizeStr)
			bytesValue := []byte{}
			for i := 0; i < 32 - size; i++ {
				bytesValue = append(bytesValue, 0)
			}
			for _, _byte := range value.([]byte) {
				bytesValue = append(bytesValue, _byte)
			}
			encValues = append(encValues, bytesValue)
		} else if strings.HasPrefix(_type, "uint") || strings.HasPrefix(_type, "int") {
			encTypes = append(encTypes, "uint256")
			encValues = append(encValues, abi.U256(value.(*big.Int)))
		}

		return encType, encValue
	}

	// Add field contents
	for _, field := range typedData.Types[primaryType] {
		_type := field["type"]
		value := data[field["name"]]

		// Structs and arrays have special handlings
		if typedData.Types[field["type"]] != nil {
			encTypes = append(encTypes, "bytes32")
			mapValue := value.(map[string]interface{})
			value = crypto.Keccak256(typedData.encodeData(field["type"], mapValue))
			encValues = append(encValues, value)
		} else if _type[len(_type)-1:] == "]" {
			encTypes = append(encTypes, "bytes32")
			parsedType := strings.Split(_type, "[")[0]
			arrayBuffer := bytes.Buffer{}
			for _, item := range value.([]interface{}) {
				if typedData.Types[parsedType] != nil {
					encoding := typedData.encodeData(parsedType, item.(map[string]interface{}))
					arrayBuffer.Write(encoding)
				} else {
					_, encValue := handlePrimitiveValue(_type, value, parsedType, item)
					arrayBuffer.Write(bytesValueOf(encValue))
				}
			}
			encValues = append(encValues, crypto.Keccak256(arrayBuffer.Bytes()))
		} else {
			encType, encValue := handlePrimitiveValue(_type, value, primaryType, data)
			encTypes = append(encTypes, encType)
			encValues = append(encValues, encValue)
		}
	}

	buffer := bytes.Buffer{}
	for _, encValue := range encValues {
		buffer.Write(bytesValueOf(encValue))
	}

	return buffer.Bytes() // https://github.com/ethereumjs/ethereumjs-abi/blob/master/lib/index.js#L336
}

func bytesValueOf(_interface interface{}) []byte {
	bytesVal, ok := _interface.([]byte)
	if ok {
		return bytesVal
	}

	switch reflect.ValueOf(_interface) {
	case reflect.ValueOf(string("")):
		return []byte(_interface.(string))
		break
	default:
		break
	}

	return []byte{}
}
