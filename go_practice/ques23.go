package main

import "fmt"

//Write a function that takes a slice of integers
//and returns true if the slice contains a given number.
//Input: ([1, 2, 3, 4, 5], 3) Output: true
func NumberCheckInSlice(slice []int, n int) bool {
	for _, value := range slice {
		if value == n {
			return true
		}
	}
	return false
}
func main() {
	fmt.Println(NumberCheckInSlice([]int{1, 2, 3, 4, 5}, 3))
}
