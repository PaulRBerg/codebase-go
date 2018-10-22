package eip712

import (
	"bytes"
	"sort"
	"unicode"

	"github.com/PaulRBerg/basics/helpers"
)

// encodeType generates the followign encoding:
// `name ‖ "(" ‖ member₁ ‖ "," ‖ member₂ ‖ "," ‖ … ‖ memberₙ ")"`
//
// each member is written as `type ‖ " " ‖ name` encodings cascade down and are sorted by name
func encodeType(_types EIP712Types) []byte {
	helpers.PrintJson("encodeType", map[string]interface{}{
		"types": _types,
	})
	//fmt.Printf("encodeType: types %v\n\n", types)

	var priorities = make(map[string]uint)
	for key := range _types {
		priorities[key] = 0
	}

	// Updates the priority for every new custom type discovered
	update := func(typeKey string, typeVal string) {
		priorities[typeVal]++

		// Importantly, we also have to check for parent types to increment them too
		for _, typeObj := range _types[typeVal] {
			_typeVal := typeObj["type"]

			firstChar := []rune(_typeVal)[0]
			if unicode.IsUpper(firstChar) {
				priorities[_typeVal]++
			}
		}
	}

	// Checks if referenced type has already been visited to optimise algo
	visited := func(arr []string, val string) bool {
		for _, elem := range arr {
			if elem == val {
				return true
			}
		}
		return false
	}

	for typeKey, typeArr := range _types {
		var typeValArr []string

		for _, typeObj := range typeArr {
			typeVal := typeObj["type"]
			//if typeKey == typeVal {
			//	panic(fmt.Errorf("type %s cannot reference itself", typeVal))
			//}

			firstChar := []rune(typeVal)[0]
			if unicode.IsUpper(firstChar) {
				if _types[typeVal] != nil {
					if !visited(typeValArr, typeVal) {
						typeValArr = append(typeValArr, typeVal)
						update(typeKey, typeVal)
					}
				}
			}
			//	} else {
			//		panic(fmt.Errorf("referenced type %s is undefined", typeVal))
			//	}
			//} else {
			//	if !isStandardType(typeVal) {
			//		if types[typeVal] != nil {
			//			panic(fmt.Errorf("Custom type %s must be capitalized", typeVal))
			//		} else {
			//			panic(fmt.Errorf("Unknown type %s", typeVal))
			//		}
			//	}
			//}
		}

		typeValArr = []string{}
	}

	sortedPriorities := sortByPriorityAndName(priorities)
	var buffer bytes.Buffer
	for _, priority := range sortedPriorities {
		typeKey := priority.Type
		typeArr := _types[typeKey]

		buffer.WriteString(typeKey)
		buffer.WriteString("(")

		for _, typeObj := range typeArr {
			buffer.WriteString(typeObj["type"])
			buffer.WriteString(" ")
			buffer.WriteString(typeObj["name"])
			buffer.WriteString(",")
		}

		buffer.Truncate(buffer.Len() - 1)
		buffer.WriteString(")")
	}

	return buffer.Bytes()
}

// Helper function to sort types by priority and name. Priority is calculated b
// based upon the number of references.
func sortByPriorityAndName(input map[string]uint) []EIP712TypePriority {
	var priorities []EIP712TypePriority
	for key, val := range input {
		priorities = append(priorities, EIP712TypePriority{key, val})
	}
	// Alphabetically
	sort.Slice(priorities, func(i, j int) bool {
		return priorities[i].Type < priorities[j].Type
	})
	// Priority
	sort.Slice(priorities, func(i, j int) bool {
		return priorities[i].Value > priorities[j].Value
	})

	//for _, priority := range priorities {
	//	fmt.Printf("%s, Priority %d\n", priority.Type, priority.Value)
	//}
	//fmt.Printf("\n")

	return priorities
}
