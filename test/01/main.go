package main

import "fmt"

func main() {
	x := sum(1, 2, 3)
	fmt.Println(x)
}

func sum(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}

	return total
}
