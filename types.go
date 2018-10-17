package main

import (
	"fmt"
	"github.com/PaulRBerg/go-ethereum/accounts/abi"
	"github.com/holiman/go-ethereum/common"
	"math/big"
)

func maintypes() {
	uint256()
}

func uint256() {
	var int64Val int64 = 3
	encodedVal := abi.U256(big.NewInt(int64Val))
	fmt.Println(common.Bytes2Hex(encodedVal))
}
