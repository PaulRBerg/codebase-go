package main

import (
	"fmt"
	"strings"
)

func mainstrings() {
	var str string = "MetaMask is an orange fox"
	var trimmedStr string = strings.TrimPrefix(str, "MetaMask ")
	fmt.Println(trimmedStr)

	str = "bytes24"
	trimmedStr = strings.TrimPrefix(str, "bytes")
	fmt.Println(trimmedStr)

	str = "Rob: Hello everyone!"
	user := str[:strings.IndexByte(str, ':')]
	fmt.Println(user)

	str = "SomeType[50]"
	typeStr := str[:strings.IndexByte(str, '[')]
	fmt.Println(typeStr)

	str = "Rob: Hello everyone!"
	parts := strings.Split(str, ":")
	fmt.Printf("%q\n", parts)
	if len(parts) > 1 {
		fmt.Println("User:", parts[0])
	}
}
