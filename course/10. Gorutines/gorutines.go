package main

import (
	"fmt"
	"time"
)

func sayHi(i int) {
	fmt.Println("Hello, world! I'm", i, "goroutine!")
}

func test_gorutines() {
	for i := range 5 {
		go sayHi(i)
	}
	time.Sleep(1 * time.Second) // Даем время горутинам на выполнение
}
