package main

import (
	"fmt"
)

func main() {
	channel := make(chan int)
	quit := make(chan int)

	go receiveQuit(channel, quit)
	sendToChannel(channel, quit)
}

func receiveQuit(c chan int, q chan int) {
	for i := 0; i < 50; i++ {
		fmt.Println("Received:", <-c)
	}
	q <- 0
}

func sendToChannel(c chan int, q chan int) {
	anything := 1
	for {
		select {
		case c <- anything:
			anything++
		case <-q:
			return
		}
	}
}
