package main

import (
	"fmt"
	"sort"
)

// Write a function that takes a slice of integers and returns the second-largest number in the slice.
// Input: [1, 2, 3, 4, 5] Output: 4
func SecondLargestNumber(slice []int) int {
	//sorting the slice in ascending order
	sort.Ints(slice)
	//returning the second largest
	return slice[len(slice)-2]
}

func main() {
	fmt.Println(SecondLargestNumber([]int{1, 2, 3, 4, 5}))
}
