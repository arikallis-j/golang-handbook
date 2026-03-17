package main

import "fmt"

func test_maps() {
	m := make(map[string]string) // создаём map — о применении функции make для создания переменных типа map будет рассказано ниже
	m["foo"] = "bar"             // заполняем парами «ключ-значение»
	m["ping"] = "pong"
	fmt.Println(m) // печатаем

	var m1 = map[string]string{}
	m1["test"] = "1"
	fmt.Println(m1)

}
