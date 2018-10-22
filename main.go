package main

import (
	"fmt"

	"github.com/PaulRBerg/basics/eip712"
)

func main() {
	eip712.MainEncode712V2()
}

func ForLoops() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
