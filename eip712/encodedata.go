package eip712

import (
	"bytes"
	"fmt"
	"github.com/PaulRBerg/basics/helpers"
	"github.com/PaulRBerg/go-ethereum/accounts/abi"
	"github.com/PaulRBerg/go-ethereum/common"
	"github.com/PaulRBerg/go-ethereum/crypto"
	"math/big"
	"math/rand"
	"reflect"
	"time"
)

//type EIP712Message map[string]interface{}

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

// `enc(value₁) ‖ enc(value₂) ‖ … ‖ enc(valueₙ)`
//
// each encoded member is 32-byte long
func encodeData(types map[string]EIP712Type, key string, val interface{}, depth int) string {
	helpers.PrintJson("hashStruct", map[string]interface{}{
		"key":   key,
		"val":   val,
		"depth": depth,
	})
	//fmt.Printf("encodeData: \"key\" %s, \"val\" %v, \"depth\" %d\n\n", key, val, depth)

	var buffer bytes.Buffer

	switch val.(type) {
	case EIP712Data:
		for mapKey, mapVal := range val.(EIP712Data) {
			if reflect.TypeOf(mapVal) == reflect.TypeOf(EIP712Data{}) {
				hash := hashStruct(types, mapKey, mapVal.(EIP712Data), depth+1)
				buffer.WriteString(hash.String())
			} else {
				str := encodeData(types, mapKey, mapVal, depth+1)
				buffer.WriteString(str)
			}
		}
		break

	case bool:
		boolVal, _ := val.(bool)
		var int64Val int64
		if boolVal {
			int64Val = 1
		}
		encodedVal := abi.U256(big.NewInt(int64Val))
		fmt.Printf("bool encoded value:", encodedVal)
		buffer.Write(encodedVal)
		break

	case string:
		bytesVal := common.FromHex(val.(string))
		hash := common.BytesToHash(crypto.Keccak256(bytesVal))
		buffer.WriteString(hash.String())
		break

	default:
		arr := [...]string{"(sarah)", "(sophie)", "(ivana)"}
		rand.Seed(time.Now().UnixNano())
		buffer.WriteString(arr[rand.Intn(3)])
		break
	}

	return buffer.String()
}
