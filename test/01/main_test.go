package main

import "testing"

func TestSum(t *testing.T) {
	test := sum(3, 2, 1)
	result := 6
	if test != result {
		t.Error("Expected:", result, "Got:", test)
	}
}

// func TestSumError(t *testing.T) {
// 	test := sum(3, 2, 1)
// 	result := 7
// 	if test != result {
// 		t.Error("Expected:", result, "Got:", test)
// 	}
// }

func TestMultiply(t *testing.T) {
	test := multiply(10, 10)
	result := 100
	if test != result {
		t.Error("Expected:", result, "Got:", test)
	}
}

//Testes files with _test
//Function test starts With Test
