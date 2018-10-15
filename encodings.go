package main

import (
	"bytes"
	"fmt"
	"github.com/PaulRBerg/go-ethereum/common"
	"github.com/PaulRBerg/go-ethereum/crypto"
	"sort"
	"unicode"
)

type EIP712Types map[string][]map[string]string

type EIP712SortedPriority struct {
	Type	string
	Value 	uint
}

var typesA = map[string][]map[string]string{
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

var typesB = map[string][]map[string]string{
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

func mainencodings() {
	_, err := encodeType(typesB, "Mail")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	//fmt.Printf("hash: %s", hash.String())
}

func encodeType(types EIP712Types, primaryType string) (common.Hash, error) {
	var priorities = make(map[string]uint)
	for key := range types { priorities[key] = 0 }

	visited := func(arr []string, val string) bool {
		for _, elem := range arr {
			if elem == val {
				return true
			}
		}
		return false
	}

	for _, typeArr := range types {
		var typeValArr []string

		for _, typeObj := range typeArr {
			typeVal := typeObj["type"]
			firstChar := []rune(typeVal)[0]

			if unicode.IsUpper(firstChar) {
				if (types[typeVal] != nil) {
					if (!visited(typeValArr, typeVal)) {
						typeValArr = append(typeValArr, typeVal)
						priorities[typeVal]++
					}
				} else {
					return common.Hash{}, fmt.Errorf("referenced type %s is undefined", typeVal)
				}
			} else {
				if !isStandardType(typeVal) {
					if types[typeVal] != nil {
						return common.Hash{}, fmt.Errorf("Custom type %s must be capitalized", typeVal)
					} else {
						return common.Hash{}, fmt.Errorf("Unknown type %s", typeVal)
					}
				}
			}
		}

		typeValArr = []string{}
	}


	sortedPriorities := types.SortByPriority(priorities)
	for _, kv := range sortedPriorities {
		fmt.Printf("%s, %d\n", kv.Type, kv.Value)
	}

	var buffer bytes.Buffer
	for _, priority := range sortedPriorities {
		typeKey := priority.Type
		typeArr := types[typeKey]

		buffer.WriteString(typeKey)
		buffer.WriteString("(")

		for _, typeObj := range typeArr {
			buffer.WriteString(typeObj["type"])
			buffer.WriteString(" ")
			buffer.WriteString(typeObj["name"])
			buffer.WriteString(",")
		}

		buffer.Truncate(buffer.Len()-1)
		buffer.WriteString(")")
	}

	fmt.Println("encoding:", buffer.String())

	return common.BytesToHash(crypto.Keccak256(buffer.Bytes())), nil
}

func (types *EIP712Types) SortByPriority(priorities map[string]uint) []EIP712SortedPriority {
	var sortedPriorities []EIP712SortedPriority
	for key, val := range priorities {
		sortedPriorities = append(sortedPriorities, EIP712SortedPriority{key, val})
	}
	sort.Slice(sortedPriorities, func(i, j int) bool {
		return sortedPriorities[i].Value < sortedPriorities[j].Value
	})
	return sortedPriorities
}

func isStandardType(typeStr string) bool {
	switch typeStr {
	case
		"array",
		"address",
		"boolean",
		"bytes",
		"string",
		"struct",
		"uint":
		return true
	}
	return false
}