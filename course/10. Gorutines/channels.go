package main

import "fmt"

func sayHello(i int, message chan string) {
	message <- fmt.Sprint("Hello, world! I'm ", i, " goroutine!") // Отправляем сообщение в канал
}

func main() {
	message := make(chan string) // Создаем канал для строк

	for i := range 5 {
		go sayHello(i, message)
	}

	for range 5 {
		fmt.Println(<-message) // Читаем сообщения из канала
	}
	close(message)
}
