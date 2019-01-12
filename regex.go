package main

import (
	"fmt"
	"regexp"
)

func mainregex() {
	TypeRegex()
}

func BasicRegex() {
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

func TypeRegex() {
	_ = []string{
		"bytes",
		"int",
		"uint",
	}
	r, _ := regexp.Compile(`(bytes|int|uint)(\d+)\b`)
	fmt.Println(r.MatchString("bytes256z"))
	fmt.Println(r.FindString("bytes256z"))
}
