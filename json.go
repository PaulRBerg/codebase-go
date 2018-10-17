package main

import (
	"encoding/json"
	"fmt"
	"github.com/PaulRBerg/go-ethereum/common"
	"github.com/PaulRBerg/go-ethereum/common/hexutil"
	"math/big"
)

type Bird struct {
	Species     string `json:"species"`
	Description string `json:"description"`
}

func mainjson() {
	birdJson := `{"species": "pigeon","description": "likes to perch on rocks"}`
	var bird Bird
	json.Unmarshal([]byte(birdJson), &bird)
	fmt.Printf("Species: %s, Description: %s\n", bird.Species, bird.Description)
	//Species: pigeon, Description: likes to perch on rocks

	obj, _ := json.Marshal("01020304\n")
	fmt.Println(obj)

	fmt.Println([]byte("0xdeadbeef\n"))
}

type DomainSeparator struct {
	Name    string         `json:"name"`
	Version string         `json:"version"`
	ChainId *big.Int       `json:"chainId"`
	Address common.Address `json:"address"`
	Salt    hexutil.Bytes  `json:"salt"`
}
