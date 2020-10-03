package main

import "fmt"

func main() {
	c := make(chan int)

	go myLoop(10, c)

	for v := range c {
		fmt.Println("Received from channel:", v)
	}
}

func myLoop(t int, s chan<- int) {
	for i := 0; i < t; i++ {
		s <- i
	}
	close(s)
}
