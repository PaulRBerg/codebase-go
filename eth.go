package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

func maineth() {
	PrintEthereumAddress()
}

func PrintByte() {
	const prependedSigByte = 0x19
	const byteVersion = 0x45
	const data = "0xdeadbeef"

	msg := fmt.Sprintf("\\x%x\\x%xEthereum Signed Message", prependedSigByte, byteVersion)
	fmt.Println(msg)
}

func PrintEthereumAddress() {
	fmt.Println(common.Address{}.String())
}
