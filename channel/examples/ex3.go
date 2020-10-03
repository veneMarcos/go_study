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

// - Assignment/conversion:
//     - de geral para específico
//     - de específico para geral não
//     - Exemplos:
//         - geral pra específico: https://play.golang.org/p/H1uk4YGMBB
//         - específico pra específico: https://play.golang.org/p/8JkOnEi7-a
//         - específico pra geral: https://play.golang.org/p/4sOKuQRHq7
//         - atribuição tipos !=: https://play.golang.org/p/bG7H6l03VQ
