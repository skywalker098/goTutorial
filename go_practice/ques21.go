package main

import "fmt"

//Write a function that takes a slice of integers and returns a new slice with the elements in alternating order.
//Input: [1, 2, 3, 4, 5] Output: [1, 3, 5, 2, 4]
func AlternateSlice(slice []int) []int {
	res := []int{}
	for pos, value := range slice {
		if pos == 0 {
			res = append(res, value)
		} else if pos%2 == 0 {
			res = append(res, value)
		}
	}
	return res
}

func main() {
	fmt.Println(AlternateSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}
