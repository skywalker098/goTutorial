package main

import "fmt"

//Write a function that takes a slice of integers and returns the largest number in the slice. Input: [1, 2, 3, 4, 5] Output: 5
func LargestNumber(slice []int) int {
	largestNumber := slice[0]
	// for i := 1; i <= len(slice); i++ {
	// 	if largestNumber < slice[i] {
	// 		largestNumber = slice[i]
	// 	}

	// }
	for _, v := range slice {
		if largestNumber < v {
			largestNumber = v
		}
	}
	return largestNumber
}

func main() {
	fmt.Println(LargestNumber([]int{10, 1, 2, 3, 6, 5, 3, 4}))
}
