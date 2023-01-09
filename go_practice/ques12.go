package main

import "fmt"

//Write a function that takes a slice of integers and returns the index of the first occurrence of a given number.
//Input: ([1, 2, 3, 2, 4, 5], 2) Output: 1

// func IndexFirstOccurence(slice []int, x int) int {
// 	var pos int
// 	for i := 0; i < len(slice); i++ {
// 		if slice[i] == x {
// 			pos = i
// 			break
// 		}
// 	}
// 	return pos
// }
func IndexFirstOccurence(slice []int, x int) int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == x {
			//without using another variable to store i
			return i
			break
		}
	}
	return -1
}

func main() {
	fmt.Println(IndexFirstOccurence([]int{1, 2, 3, 4, 5}, 2))
}
