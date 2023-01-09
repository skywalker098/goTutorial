package main

import "fmt"

//Write a function that takes a slice of integers and returns the sum of all the elements in the slice.
//Input: [1, 2, 3, 4, 5] Output: 15
func SumSlice(slice []int) int {
	//var sum int = 0
	total := 0
	// for i := 0; i < len(slice); i++ {
	// 	sum = sum + slice[i]
	// }
	// return sum
	for _, v := range slice {
		total += v
	}
	return total
}

func main() {
	fmt.Println(SumSlice([]int{1, 2, 3, 4, 5}))
}
