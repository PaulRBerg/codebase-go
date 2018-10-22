package helpers

import (
	"encoding/json"
	"fmt"
)

func PrintJson(label string, output map[string]interface{}) {
	jsonVal, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s:", label)
	fmt.Print(string(jsonVal))
	fmt.Print("\n\n")
}
