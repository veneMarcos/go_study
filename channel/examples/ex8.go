package main

import (
	"fmt"
	"sync"
)

func main() {
	par := make(chan int)
	odd := make(chan int)
	converge := make(chan int)

	go send(par, odd)
	go receive(par, odd, converge)

	for v := range converge {
		fmt.Println("received value:", v)
	}
}

func send(p, o chan int) {
	x := 100
	for i := 0; i < x; i++ {
		if i%2 == 0 {
			p <- i
		} else {
			o <- i
		}
	}
	close(p)
	close(o)
}

func receive(p, i, c chan int) {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		for v := range p {
			c <- v
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for v := range i {
			c <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(c)
}
