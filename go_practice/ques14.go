package main

import "fmt"

//Write a function that takes a slice of integers
//and returns a new slice with the elements in reverse order.
//Input: [1, 2, 3, 4, 5] Output: [5, 4, 3, 2, 1]
func ReverseSlice(slice []int) []int {
	tempSlice := []int{}
	sliceLength := len(slice)
	for i := sliceLength; i >= 0; i-- {
		tempSlice = append(tempSlice, slice[i])
	}
	return tempSlice
}

func main() {
	fmt.Println(ReverseSlice([]int{1, 2, 3, 4, 5}))
}
