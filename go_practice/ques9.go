package main

import "fmt"

//Write a function that takes a slice of integers and returns the average of all the elements in the slice.
//Input: [1, 2, 3, 4, 5] Output: 3
func SliceAvg(slice []int) int {
	total := 0
	sliceLength := len(slice)
	for _, v := range slice {
		total += v
	}
	//avg := total / len(slice)
	return total / sliceLength
}

func main() {
	fmt.Println(SliceAvg([]int{1, 2, 3, 4, 5, 6}))
}
