package main

import (
	"fmt"
	"time"
)

type Token struct {
	Data      string
	Recipient int
	TTL       int
}

func node(id int, in <-chan Token, out chan<- Token) {
	for {
		token := <-in
		if token.TTL <= 0 {
			fmt.Println("TTL закончился на номере ", id)
			continue
		}

		if id == token.Recipient {
			fmt.Printf("Узел %d получил сообщение: %s\n", id, token.Data)
			continue
		}

		token.TTL--
		out <- token
	}
}
func main() {
	var N int
	fmt.Println("Введите количество узлов: ")
	fmt.Scanln(&N)
	channels := make([]chan Token, N)
	for i := range channels {
		channels[i] = make(chan Token)
	}

	for i := 0; i < N; i++ {
		next := (i + 1) % N
		go node(i, channels[i], channels[next])
	}

	var message string
	message = "Привет, мир"
	fmt.Println("Введите номер получателя от 0 до", N-1)
	var recipient int
	fmt.Scanln(&recipient)
	if recipient > N-1 {
		recipient = N - 1
		fmt.Println("recipient изменен на ", N-1)
	}
	fmt.Println("Введите TTL:")
	var ttl int
	fmt.Scanln(&ttl)
	initialToken := Token{Data: message, Recipient: recipient, TTL: ttl}
	channels[0] <- initialToken
	time.Sleep(time.Second * 5)
}
