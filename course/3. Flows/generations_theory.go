package main

import (
	"fmt"
	"time"
)

var year = 1940

func test_generation() {
	date := time.Date(1004, 2, 3, 0, 0, 0, 0, time.UTC)

	switch year := date.Year(); {
	case year >= 1946 && year <= 1964:
		fmt.Println("Привет, бумер!")
	case year >= 1965 && year <= 1980:
		fmt.Println("Привет, представитель X!")
	case year >= 1981 && year <= 1996:
		fmt.Println("Привет, миллениал!")
	case year >= 1997 && year <= 2012:
		fmt.Println("Привет, зумер!")
	case year >= 2013:
		fmt.Println("Привет, альфа!")
	default:
		fmt.Println("Ты стар, просто super star!")
	}
}
