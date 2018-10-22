package eip712

import (
	"bytes"
	"errors"
	"github.com/PaulRBerg/go-ethereum/common"
	"strings"

	"github.com/PaulRBerg/go-ethereum/crypto"
)

func (typedData *TypedData) typeHash(primaryType string) []byte {
	return crypto.Keccak256(typedData.encodeTypeV2(primaryType))
}

// encodeData generates the following encoding:
// `enc(value₁) ‖ enc(value₂) ‖ … ‖ enc(valueₙ)`
//
// each encoded member is 32-byte long
func (typedData *TypedData) encodeDataV2(primaryType string, data map[string]interface{}) []byte {
	encTypes := []string{}
	encValues := []interface{}{}

	// Add typehash
	encTypes = append(encTypes, "bytes32")
	encValues = append(encValues, typedData.typeHash(primaryType))

	// Add field contents
	for _, field := range typedData.Types[primaryType] {
		_type := field["type"]
		value := data[field["name"]]
		if typedData.Types[field["type"]] != nil {
			encTypes = append(encTypes, "bytes32")
			mapValue := value.(map[string]interface{})
			value = crypto.Keccak256(typedData.encodeDataV2(field["type"], mapValue))
			encValues = append(encValues, value)
		} else if _type == "address" {
			encTypes = append(encTypes, "address")
			addressValue := common.HexToAddress(value.(string))

			bytesValue := []byte{}
			for i := 0; i < 12; i++ {
				bytesValue = append(bytesValue, 0)
			}
			for _, _byte := range addressValue.Bytes() {
				bytesValue = append(bytesValue, _byte)
			}

			encValues = append(encValues, bytesValue) // hopefully this means uint160 encoding?
		} else if _type == "string" || _type == "bytes" {
			encTypes = append(encTypes, "bytes32")
			value := crypto.Keccak256([]byte(value.(string)))
			encValues = append(encValues, value)
		} else if strings.Contains(_type, "]") {
			panic(errors.New("TODO: Arrays currently unimplemented in encodeData"))
		} else {
			encTypes = append(encTypes, field["type"])
			encValues = append(encValues, []byte(value.(string)))
		}
	}

	buffer := bytes.Buffer{}
	for _, encValue := range encValues {
		if stringValue, ok := encValue.(string); ok {
			buffer.WriteString(stringValue)
		} else {
			buffer.Write(encValue.([]byte))
		}
	}

	return buffer.Bytes() // https://github.com/ethereumjs/ethereumjs-abi/blob/master/lib/index.js#L336
}