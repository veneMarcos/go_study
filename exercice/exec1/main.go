package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	newgoroutines(200)
	wg.Wait()
}

func newgoroutines(i int) {
	wg.Add(i)
	for j := 1; j <= i; j++ {
		x := j
		go func(i int) {
			fmt.Println("Goroutines number: ", i)
			wg.Done()
		}(x)
	}
}
