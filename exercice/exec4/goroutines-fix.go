package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex
var count int

const goroutinesQuantity = 100

func main() {
	fmt.Println("CPUs ", runtime.NumCPU())

	createGoroutines(goroutinesQuantity)
	wg.Wait()

	fmt.Println("Goroutines: ", goroutinesQuantity)
	fmt.Println("Count:", count)
}

func createGoroutines(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func() {
			mu.Lock()
			v := count
			runtime.Gosched()
			v++
			count = v
			mu.Unlock()
			wg.Done()
		}()
	}
}

// You can running your code with -race, this is make a analys of your code.
// go run -race goroutines-with-racedondition.go
