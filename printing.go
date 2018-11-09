package main

import "fmt"

func mainprinting() {
	bytes := []byte("0xdeadbeef")
	fmt.Printf("%s\n", string(bytes)) // bytes or string(bytes) is the same for fmt.Printf

	fmt.Printf("%x\n", bytes) // 30786465616462656566
	fmt.Printf("%s\n", bytes) // 0xdeadbeef
}

func PrettyPrint() {

}
