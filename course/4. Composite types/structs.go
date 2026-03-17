package main

import "fmt"

type Person struct {
	Name string // Имя
	Age  int    // Возраст
}

func test_struct() {
	man := Person{"Alex", 30}
	fmt.Println(man)
	fmt.Printf("Man %#v", man)
}
