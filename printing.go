package main

import (
	"encoding/json"
	"fmt"
)

func mainprinting() {
	PrettyPrint()
}

func PrintBytes() {
	bytes := []byte("0xdeadbeef")
	fmt.Printf("%s\n", string(bytes)) // bytes or string(bytes) is the same for fmt.Printf

	fmt.Printf("%x\n", bytes) // 30786465616462656566
	fmt.Printf("%s\n", bytes) // 0xdeadbeef
}

func PrettyPrint() {
	fancyMap := map[string]interface{}{
		"foo": "bar",
		"baz": "waldo",
	}
	_ = map[string]interface{}{
		"dandy": "stuff",
	}

	jsonOutput, _ := json.MarshalIndent(fancyMap, "", "\t")
	fmt.Printf("%s\n", jsonOutput)
}

func PrettyPrintV2() {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.MarshalIndent(group, "", "\t")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%s\n", b)
}