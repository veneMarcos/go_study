//assignment/conversion
package main

import "fmt"

func main() {
	channel := make(chan int)    //bidirecional
	channelR := make(<-chan int) // receive
	channelS := make(chan<- int) // send

	fmt.Println("-----")
	fmt.Printf("channel\t%T\tbidirecional\n", channel)
	fmt.Printf("channelR\t%T\treceive\n", channelR)
	fmt.Printf("channelS\t%T\tsend\n", channelS)

	// general to specific converts
	fmt.Println("-----")
	fmt.Printf("channel\t%T\n", (<-chan int)(channel))
	fmt.Printf("channel\t%T\n", (chan<- int)(channel))
}
