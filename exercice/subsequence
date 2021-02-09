package main

func IsValidSubsequence(array []int, sequence []int) bool {
	check := 0
	
	for i, _ := range array {
		if array[i] == sequence[check] {
			check++
		}
		
		if check == len(sequence) {
			return true
		}		
	}
	
	return false
}
