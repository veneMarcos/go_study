package main

import "fmt"

func main() {
	x := sum(1, 2, 3)
	y := multiply(10, 10)
	fmt.Println(x, y)
}

func sum(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}

	return total
}

func multiply(i ...int) int {
	total := 1 // starts 1 because multiply por 0 is 0
	for _, v := range i {
		total = total * v
	}

	return total
}
