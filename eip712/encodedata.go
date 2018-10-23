package eip712

import (
	"bytes"
	"math/big"
	"reflect"
	"strconv"
	"strings"

	"github.com/PaulRBerg/go-ethereum/accounts/abi"
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
	handlePrimitiveValue := func(encType string, encValue interface{}, primaryType string, data interface{}) (string, interface{}) {
		var primitiveEncType string
		var primitiveEncValue interface{}

		switch encType {
		case "address":
			primitiveEncType = "address"
			bytesValue := []byte{}
			for i := 0; i < 12; i++ {
				bytesValue = append(bytesValue, 0)
			}
			for _, _byte := range encValue.(common.Address).Bytes() {
				bytesValue = append(bytesValue, _byte)
			}
			primitiveEncValue = bytesValue
			break
		case "bool":
			primitiveEncType = "uint256"
			var int64Val int64
			if encValue.(bool) {
				int64Val = 1
			}
			primitiveEncValue = abi.U256(big.NewInt(int64Val))
			break
		case "bytes", "string":
			primitiveEncType = "bytes32"
			value := crypto.Keccak256(bytesValueOf(encValue))
			primitiveEncValue = value
			break
		default:
			if strings.HasPrefix(encType, "bytes") {
				encTypes = append(encTypes, "bytes32")
				sizeStr := strings.TrimPrefix(encType, "bytes")
				size, _ := strconv.Atoi(sizeStr)
				bytesValue := []byte{}
				for i := 0; i < 32-size; i++ {
					bytesValue = append(bytesValue, 0)
				}
				for _, _byte := range encValue.([]byte) {
					bytesValue = append(bytesValue, _byte)
				}
				encValues = append(encValues, bytesValue)
			} else if strings.HasPrefix(encType, "uint") || strings.HasPrefix(encType, "int") {
				encTypes = append(encTypes, "uint256")
				encValues = append(encValues, abi.U256(encValue.(*big.Int)))
			}
			break
		}
		return primitiveEncType, primitiveEncValue
	}

	// Add field contents. Structs and arrays have special handlings.
	for _, field := range typedData.Types[primaryType] {
		encType := field["type"]
		encValue := data[field["name"]]

		if encType[len(encType)-1:] == "]" {
			encTypes = append(encTypes, "bytes32")
			parsedType := strings.Split(encType, "[")[0]
			arrayBuffer := bytes.Buffer{}
			for _, item := range encValue.([]interface{}) {
				if typedData.Types[parsedType] != nil {
					encoding := typedData.encodeData(parsedType, item.(map[string]interface{}))
					arrayBuffer.Write(encoding)
				} else {
					_, encValue := handlePrimitiveValue(encType, encValue, parsedType, item)
					arrayBuffer.Write(bytesValueOf(encValue))
				}
			}
			encValues = append(encValues, crypto.Keccak256(arrayBuffer.Bytes()))
		} else if typedData.Types[field["type"]] != nil {
			encTypes = append(encTypes, "bytes32")
			mapValue := encValue.(map[string]interface{})
			encValue = crypto.Keccak256(typedData.encodeData(field["type"], mapValue))
			encValues = append(encValues, encValue)
		} else {
			encType, encValue := handlePrimitiveValue(encType, encValue, primaryType, data)
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
