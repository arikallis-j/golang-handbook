package main

import "fmt"

func test_v() {
	var a, b int
	a, b = 5, 10 // значения присваиваются по порядку: a == 5 и b == 10
	fmt.Println(a, b)
	a, b = b, a // swap: a == 10, b == 5
	fmt.Println(a, b)

	var intVar = 10 // переменной intVar компилятор присвоит тип int
	fmt.Println(intVar)

	var ver = "v0.0.1"
	var id int
	pi := 3.1415
	fmt.Println("ver =", ver, "id =", id, "pi =", pi)
	// ver = v0.0.1 id = 0 pi = 3.1415

	var v = 5*3 - 16
	fmt.Println(v)
}
