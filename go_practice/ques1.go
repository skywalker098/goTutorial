package main

import "fmt"

//Write a function that takes a slice of integers and returns a new slice with only the even numbers.
//Input: [1, 2, 3, 4, 5] Output: [2, 4]
func EvenSlice(slice []int) []int {
	newSlice := []int{}
	for _, v := range slice {
		if v%2 == 0 {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

func main() {
	fmt.Println(EvenSlice([]int{1, 2, 3, 4, 5, 6}))
}
