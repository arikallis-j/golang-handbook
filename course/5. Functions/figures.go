package main

import (
	"fmt"
	"math"
)

type figures int

const (
	square   figures = iota // квадрат
	circle                  // круг
	triangle                // равносторонний треугольник
)

func area(fig figures) (func(float64) float64, bool) {
	var ar func(float64) float64
	var ok bool

	if fig == square {
		ar, ok = func(a float64) float64 { return a * a }, true
	} else if fig == circle {
		ar, ok = func(r float64) float64 { return math.Pi * r * r }, true
	} else if fig == triangle {
		ar, ok = func(a float64) float64 { return math.Sqrt(3) / 4 * a * a }, true
	}
	return ar, ok
}

func test_funcs() {
	var test figures = 5
	myFigure := test

	ar, ok := area(myFigure)
	if !ok {
		fmt.Println("Ошибка")
		return
	}
	x := 1.0

	myArea := ar(x)
	fmt.Println(myArea)
}
