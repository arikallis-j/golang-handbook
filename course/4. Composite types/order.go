package main

import (
	"fmt"
)

func test_order() {
	var price = map[string]int{
		"хлеб":     50,
		"молоко":   100,
		"масло":    200,
		"колбаса":  500,
		"соль":     20,
		"огурцы":   200,
		"сыр":      600,
		"ветчина":  700,
		"буженина": 900,
		"помидоры": 250,
		"рыба":     300,
		"хамон":    1500,
	}

	for k, v := range price {
		if v > 500 {
			fmt.Println(k)
		}
	}

	order := []string{"хлеб", "буженина", "сыр", "огурцы"}

	var cost int

	for _, v := range order {
		cost += price[v]
	}
	fmt.Println(cost)
}
