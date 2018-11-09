package main

import (
	"fmt"
)

func main() {
	mainprinting()
	//eip712.MainEncode712()
}

func ForLoops() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
