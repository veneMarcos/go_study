package main

import "fmt"

func main() {
	channel := make(chan int)

	// As funcoes com canais devem rodar de maneira concorrentes então uma delas deve estar em uma goroutines
	go send(channel)

	receiver(channel)
}

//Direcionamento de canais

func send(s chan<- int) {
	s <- 42
}

func receiver(r <-chan int) {
	fmt.Println("The value received on channel is", <-r)
}
