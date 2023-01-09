package main

import (
	"fmt"
)

//Write a function that takes a slice of integers
//and returns true if the slice is sorted in descending order.
//Input: [5, 4, 3, 2, 1] Output: true

func DescendingOrderCheck(slice []int) bool {
	for i := 0; i < len(slice)-1; i++ {
		if slice[i+1] > slice[i] {
			return false
		}
	}
	return true

}

func main() {
	fmt.Println(DescendingOrderCheck([]int{5, 4, 3, 2, 1, 4}))
}
