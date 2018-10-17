package main

import (
	"encoding/hex"
	"fmt"
	"github.com/PaulRBerg/go-ethereum/crypto"
)

func maincrypto() {
	sighash, msg := SignDataPlain("0xdeadbeef")
	fmt.Println("sighash", hex.EncodeToString(sighash))
	fmt.Println("msg:", msg)
}

func SignDataPlain(data string) ([]byte, string) {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg)), msg
}
