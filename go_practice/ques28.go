package main

import (
	"fmt"
	"sort"
)

//Write a function that takes a slice of integers
//and returns the smallest number that is divisible by all elements in the slice.
// Input: [2, 3, 5] Output: 30

func LCM(slice []int) []int {
	//tempSlice := []int{}
	sort.Ints(slice)
	min := slice[0]
	return slice
}

func main() {
	fmt.Println(LCM([]int{2, 5, 3}))
}
