package main

import "fmt"

func main() {
	par := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go sendNumbers(par, odd, quit)
	receive(par, odd, quit)
}

func sendNumbers(par, odd chan int, quit chan bool) {
	total := 100
	for i := 0; i < total; i++ {
		if i%2 == 0 {
			par <- i
		} else {
			odd <- i
		}
	}
	close(par)
	close(odd)
	quit <- true
}

func receive(par, odd chan int, quit chan bool) {
	for {
		select {
		case v := <-par:
			fmt.Println("The received number:", v, "is par")
		case v := <-odd:
			fmt.Println("The received number:", v, "is odd")
		case v, ok := <-quit:
			if !ok {
				fmt.Println("Houston we have a problem", v)
			} else {
				fmt.Println("Finish. Received:", v)
			}
			return
		}
	}
}
