package main

import "fmt"

//Write a function that takes a slice of integers and a number
//and returns the number of times the number appears in the slice.
//Input: ([1, 2, 3, 2, 4, 5], 2) Output: 2
func FrequencyCounter(slice []int, x int) int {
	temp := 0
	for _, v := range slice {
		if v == x {
			temp += 1
		}
	}
	return temp
}

func main() {
	fmt.Println(FrequencyCounter([]int{1, 2, 3, 4, 5, 2}, 2))
}
