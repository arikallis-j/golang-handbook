package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x *bool
	y := reflect.ValueOf(x)
	fmt.Println(y.Kind())
	fmt.Println(y.Type())
}
