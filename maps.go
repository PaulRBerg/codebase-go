package main

import "fmt"

func mainmaps() {
	cities := make(map[uint32]string)
	cities[0] = "San Francisco"
	cities[1] = "London"
	cities[2] = "Iasi"

	fmt.Println(cities)
}
