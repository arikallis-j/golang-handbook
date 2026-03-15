package main

const (
	one = 2*iota + 1 // укажите здесь формулу с iota
	three
	five
	seven
	_
	eleven
)

func test_iota() {
	println(one, three, five, seven, eleven)
}
