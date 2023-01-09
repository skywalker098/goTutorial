package main

import "fmt"

//Write a function that takes a slice of integers
//and returns a new slice with all elements that are divisible by a given number.
//Input: ([1, 2, 3, 4, 5], 2) Output: [2, 4]
func DivisibilityByNumber(slice []int, n int) []int {
	res := []int{}
	for _, v := range slice {
		if v%n == 0 {
			res = append(res, v)
		}
	}
	return res
}

func main() {
	fmt.Println(DivisibilityByNumber([]int{1, 2, 3, 4, 5}, 2))
}
