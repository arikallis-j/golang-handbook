package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Person_1 struct {
	Name        string    `json:"Имя"`
	Email       string    `json:"Почта"`
	DateOfBirth time.Time `json:"-"`
}

func test_tags() {
	p := Person_1{Name: "Aлекс", Email: "alex@yandex.ru"}
	str, _ := json.Marshal(p)
	fmt.Println(string(str))
}
