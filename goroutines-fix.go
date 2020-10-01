package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("CPUs ", runtime.NumCPU())
	fmt.Println("Goroutines: ", runtime.NumGoroutine())

	vCount := 0
	totalgoroutines := 15

	var wg sync.WaitGroup
	wg.Add(totalgoroutines)

	var mu sync.Mutex

	for i := 0; i < totalgoroutines; i++ {
		go func() {
			mu.Lock()
			v := vCount
			runtime.Gosched()
			v++
			vCount = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines: ", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	fmt.Println("Valor final:", vCount)
}

// You can running your code with -race, this is make a analys of your code.
// go run -race goroutines-with-racedondition.go
