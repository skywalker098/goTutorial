package main

import (
	"fmt"
)

// Write a function that takes a slice of integers and returns true if the slice is sorted in ascending order.
// Input: [1, 2, 3, 4, 5] Output: true
func CheckSortAsc(slice []int) bool {
	// tempSlice := sort.Ints(slice)
	// if slice == tempSlice {
	// 	return true
	// }
	// return false

	for i := 0; i < (len(slice))-1; i++ {
		if slice[i] < slice[i+1] {
			continue
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(CheckSortAsc([]int{1, 2, 3, 7, 4, 5}))
}
