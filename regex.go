package main

import (
	"fmt"
	"regexp"
)

func mainregex() {
	fmt.Println("peach punch")
	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r.MatchString("peach"))
	fmt.Println(r.FindString("peach punch"))
	fmt.Println()

	// Works differently from regex101.com
	fmt.Println("bytes24")
	rnumbers, _ := regexp.Compile("([1-9]|[12][0-9]|3[01])")
	fmt.Println(rnumbers.MatchString("bytes24"))
	fmt.Println(rnumbers.FindString("bytes24"))
	fmt.Println()

	fmt.Println("SomeType[50]")
	rarrays, _ := regexp.Compile(`(\\w+)([\\d+])`)
	fmt.Println(rarrays.MatchString("SomeType[50]"))
	fmt.Println(rnumbers.FindString("SomeType[50]"))
	fmt.Println()
}
