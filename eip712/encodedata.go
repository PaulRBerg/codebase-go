package eip712

import (
	"bytes"
	"math/big"
	"reflect"
	"strings"
	"unicode"

	"github.com/PaulRBerg/basics/helpers"
	"github.com/PaulRBerg/go-ethereum/accounts/abi"
	"github.com/PaulRBerg/go-ethereum/common"
	"github.com/PaulRBerg/go-ethereum/crypto"
)

// encodeData generates the following encoding:
// `enc(value₁) ‖ enc(value₂) ‖ … ‖ enc(valueₙ)`
//
// each encoded member is 32-byte long
func encodeData(_types EIP712Types, data interface{}, dataType string, depth int) []byte {
	helpers.PrintJson("encodeData", map[string]interface{}{
		"dataType": dataType,
		"data":     data,
		"depth":    depth,
	})

	var buffer bytes.Buffer

	// handle maps
	firstChar := []rune(dataType)[0]
	if unicode.IsUpper(firstChar) {
		for mapKey, mapVal := range data.(EIP712Data) {
			nextDataType := findNextDataType(_types, dataType, mapKey)
			if reflect.TypeOf(mapVal) == reflect.TypeOf(EIP712Data{}) {
				eip712data := mapVal.(EIP712Data)
				encoding := hashStruct(_types, eip712data, nextDataType, depth+1)
				buffer.Write(encoding)
			} else {
				encoding := encodeData(_types, mapVal, nextDataType, depth+1)
				buffer.Write(encoding)
			}
		}
		return buffer.Bytes()
	}

	// TODO regex
	// handle arrays
	if strings.Contains(dataType, "[]") {
		arrayVal := data.([]interface{})
		dataType := "TODO"

		var arrayBuffer bytes.Buffer
		for obj := range arrayVal {
			objEncoding := encodeData(_types, obj, dataType, depth+1)
			arrayBuffer.Write(objEncoding)
		}

		encoding := arrayBuffer.Bytes()
		buffer.Write(encoding)
		return buffer.Bytes()
	}

	// TODO regex
	// handle bytes
	if strings.Contains(dataType, TypeBytes) {
		bytesVal := data.([]byte)
		encoding := crypto.Keccak256(bytesVal)
		buffer.Write(encoding)
	}
	//case TypeBytes1, TypeBytes2, TypeBytes3, TypeBytes4, TypeBytes5, TypeBytes6, TypeBytes7, TypeBytes8, TypeBytes9, TypeBytes10, TypeBytes11, TypeBytes12, TypeBytes13, TypeBytes14, TypeBytes15, TypeBytes16, TypeBytes17, TypeBytes18, TypeBytes19, TypeBytes20, TypeBytes21, TypeBytes22, TypeBytes23, TypeBytes24, TypeBytes25, TypeBytes26, TypeBytes27, TypeBytes28, TypeBytes29, TypeBytes30, TypeBytes31:
	//bytesVal := data.([]byte)
	//var encodedVal [32]byte
	//for i := 0; i < len(bytesVal); i++ {
	//	encodedVal = append(encodedVal, bytesVal[i])
	//}
	//break

	// TODO regex
	// handle ints with regex
	if strings.Contains(dataType, TypeInt) {
		intVal := big.NewInt(data.(int64))
		encoding := abi.U256(intVal) // not sure if this is big endian order, but it's definitey sign extended to 256 bit because of using the U256 function
		buffer.Write(encoding)
		return buffer.Bytes()
	}

	// handle what's left
	switch dataType {
	case TypeAddress:
		addressVal, _ := data.(common.Address)
		encoding := addressVal.Bytes() // hopefully this means uint160 encoding?
		buffer.Write(encoding)
		break
	case TypeBool:
		boolVal, _ := data.(bool)
		var int64Val int64
		if boolVal {
			int64Val = 1
		}
		encoding := abi.U256(big.NewInt(int64Val))
		buffer.Write(encoding)
		break
	case TypeString:
		bytesVal := common.FromHex(data.(string))
		encoding := crypto.Keccak256(bytesVal)
		buffer.Write(encoding)
		break
	default:
		break
	}

	return buffer.Bytes()
}

func findNextDataType(_types EIP712Types, mapType string, mapKey string) string {
	eip712type := _types[mapType]

	for _, mapObj := range eip712type {
		if mapObj["name"] == mapKey {
			return mapObj["type"]
		}
	}

	return ""
}
