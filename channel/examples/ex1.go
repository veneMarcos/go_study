package main

import "fmt"

func main() {
	channel := make(chan int)

	go func() {
		channel <- 42
	}()

	fmt.Println(<-channel)
}

// Exemplos de como utilizar
// Com buffer não é muito utilizado normalmente, a melhor opçao é com goroutines
//- Block: https://play.golang.org/p/dClS7vQlYE (não roda!)
//- Go routine: https://play.golang.org/p/ZbNCwUuiPi
//- Buffer: https://play.golang.org/p/32vYvCR7qn
//- Buffer block: https://play.golang.org/p/smeW6vigAT
//- Mais buffer: https://play.golang.org/p/Pe2pcboGiA
