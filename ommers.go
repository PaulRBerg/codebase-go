package main

import (
	"fmt"
)

func mainommers() {
	fmt.Println(k(4, 10, 6))
}

func k(U uint, H uint, n uint8) bool {
	if (n == 0) {
		return false
	} else {
			if (s(U, H)) {
				return true
			} else {
				H += 1 // this is the header of P(H)
				return k(U, H, n - 1)
			}
	}
}

func s(U uint, H uint) {
	// P(H) = P(U)
	// H != U
	// U is not in B(H) uncle list

	// mock-up
	return U % 4 == 0
}