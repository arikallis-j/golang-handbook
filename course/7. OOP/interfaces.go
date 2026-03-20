package main

import (
	"fmt"
	"strings"
)

func Mul(a any, b int) any {
	switch va := a.(type) {
	case int:
		return va * b
	case string:
		return strings.Repeat(va, b)
	case fmt.Stringer:
		return strings.Repeat(va.String(), b)
	default:
		return nil
	}
}

type MyString interface {
	String() string
}

type mystring string

func (s mystring) String() string {
	return string(s)
}

func test_interfaces() {
	var str_1 MyString
	str_1 = mystring("snfj")
	fmt.Println(str_1.String())
	a := 10
	fmt.Println(Mul(str_1, a))

}
