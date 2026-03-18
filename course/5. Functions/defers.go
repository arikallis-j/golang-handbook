package main

import "fmt"

var Global = 5

func change() {
	defer func(orignal_global int) {
		Global = orignal_global
	}(Global)

	Global = 10
	fmt.Println(Global)
}

func main() {
	change()
	fmt.Println(Global)
}
