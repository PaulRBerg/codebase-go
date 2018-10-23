package main

import (
	"fmt"
	"reflect"
)

func mainreflections() {
	//Types()
	//TypesV2()
	//Interfacing()
	//Setting()
	//Structing()
}

func Types() {
	var x uint8 = 16
	var t = reflect.TypeOf(x)
	fmt.Println("type of x is", t)
	fmt.Println("kind of x.t is", t.Kind())

	var v = reflect.ValueOf(x)
	fmt.Println("value of v is", v)
	fmt.Println("type of x.v is", v.Type())
	fmt.Println("kind of x.v is", v.Kind())
	fmt.Println("uint of x.v is", v.Uint())
	fmt.Println("type of uint of x.v. is", reflect.TypeOf(v.Uint()))
}

func TypesV2() {
	str := "Cow"
	fmt.Println(reflect.TypeOf(str) == reflect.TypeOf(string("")))
}

// Interface returns v's value as an interface{}.
func Interfacing() {
	var x float64 = 3.4
	var v = reflect.ValueOf(x)

	fmt.Println("value of x is", v.String())
	y := v.Interface().(float64) // y will have type float64.
	fmt.Println("interface of value of x is:", y)
}

func Setting() {
	var x float64 = 3.4
	p := reflect.ValueOf(&x) // Note: take the address of x.
	fmt.Println("type of p", p.Type())
	fmt.Println("settability of p", p.CanSet())

	v := p.Elem()
	fmt.Println("settability of v", v.CanSet())

	v.SetFloat(7.1)
	fmt.Println(v.Interface())
	fmt.Println(x)
}

type T struct {
	A int
	B string
}

func Structing() {
	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}

	fmt.Println("t was before", t)
	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	fmt.Println("t is now", t)
}
