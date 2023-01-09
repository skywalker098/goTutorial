package main

import "fmt"

//Write a function that takes a slice of integers and returns a new slice with all duplicates removed.
//Input: [1, 2, 3, 2, 4, 5, 5] Output: [1, 3, 4]

func DupRemoved(slice []int) []int {
	tempSlice := []int{}
	res := make(map[int]int)

	for _, v := range slice {
		val, ok := res[v]
		if ok {
			res[v] = val + 1
			//appending key with single occurence
		} else {
			res[v] = 1
			tempSlice = append(tempSlice, v)
		}
	}
	return tempSlice
}

func main() {
	fmt.Println(DupRemoved([]int{1, 2, 3, 2, 4, 5, 5, 6, 6}))

}
