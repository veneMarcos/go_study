package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("CPUs ", runtime.NumCPU())
	fmt.Println("Goroutines: ", runtime.NumGoroutine())

	count := 0
	totalgoroutines := 1000

	var wg sync.WaitGroup

	wg.Add(totalgoroutines)
	for i := 0; i < totalgoroutines; i++ {
		go func() {
			v := count
			runtime.Gosched()
			v++
			count = v
			wg.Done()
		}()
		fmt.Println("Goroutines: ", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	fmt.Println("Valor final:", count)
}

// You can running your code with -race, this is make a analys of your code.
// go run -race goroutines-with-racedondition.go
