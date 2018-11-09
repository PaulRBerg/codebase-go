package eip712

import (
	"bytes"
)

// dependencies returns an array of custom types ordered by their
// hierarchical reference tree
func (typedData *TypedData) dependencies(primaryType string, found []string) []string {
	includes := func(arr []string, str string) bool {
		for _, obj := range arr {
			if obj == str {
				return true
			}
		}
		return false
	}

	if includes(found, primaryType) {
		return found
	}
	if typedData.Types[primaryType] == nil {
		return found
	}
	found = append(found, primaryType)
	for _, field := range typedData.Types[primaryType] {
		for _, dep := range typedData.dependencies(field["type"], found) {
			if !includes(found, dep) {
				found = append(found, dep)
			}
		}
	}
	return found
}

// encodeType generates the following encoding:
// `name ‖ "(" ‖ member₁ ‖ "," ‖ member₂ ‖ "," ‖ … ‖ memberₙ ")"`
//
// each member is written as `type ‖ " " ‖ name` encodings cascade down and are sorted by name
func (typedData *TypedData) encodeType(primaryType string) []byte {
	// Get dependencies primary first, then alphabetical
	deps := typedData.dependencies(primaryType, []string{})
	for i, dep := range deps {
		if dep == primaryType {
			deps = append(deps[:i], deps[i+1:]...)
			break
		}
	}
	deps = append([]string{primaryType}, deps...)

	// Format as a string with fields
	var buffer bytes.Buffer
	for _, dep := range deps {
		buffer.WriteString(dep)
		buffer.WriteString("(")
		for _, obj := range typedData.Types[dep] {
			buffer.WriteString(obj["type"])
			buffer.WriteString(" ")
			buffer.WriteString(obj["name"])
			buffer.WriteString(",")
		}
		buffer.Truncate(buffer.Len() - 1)
		buffer.WriteString(")")
	}
	return buffer.Bytes()
}
