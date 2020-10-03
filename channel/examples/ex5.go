package main

import "fmt"

func main() {
	a := make(chan int)
	b := make(chan int)
	x := 500

	go func(t int) {
		for i := 0; i < t; i++ {
			a <- i
		}
	}(x / 2)

	go func(t int) {
		for i := 0; i < t; i++ {
			b <- i
		}
	}(x / 2)

	for i := 0; i < x; i++ {
		select {
		case v := <-a:
			fmt.Println("Channel A:", v)
		case v := <-b:
			fmt.Println("Channel B:", v)
		}
	}
}
