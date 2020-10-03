package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var mu sync.Mutex
var count int32

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
			atomic.AddInt32(&count, 1)
			v := atomic.LoadInt32(&count)
			runtime.Gosched()
			fmt.Println(v)
			wg.Done()
		}()
	}
}

// You can running your code with -race, this is make a analys of your code.
// go run -race goroutines-with-racedondition.go
