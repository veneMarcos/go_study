package main

import "testing"

type test struct {
	data   []int
	answer int
}

func TestTableSum(t *testing.T) {
	tests := []test{
		{data: []int{1, 2, 3}, answer: 6},
		{[]int{10, 11, 12}, 33},
		{[]int{-5, 0, 5, 10}, 10},
	}
	for _, v := range tests {
		z := sum(v.data...)
		if z != v.answer {
			t.Error("Expected:", v.answer, "Got:", z)
		}
	}
}
