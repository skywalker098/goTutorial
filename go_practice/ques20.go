package main

import "fmt"

//Write a function that takes a slice of integers and returns a new slice with all the prime numbers.
//Input: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10] Output: [2, 3, 5, 7]

func PrimeNoSlice(slice []int) []int {
	res := []int{}
	start, flag := 0, 0

	if slice[start] == 1 {
		res = append(res, slice[start])
	}
	for _, v := range slice {
		for i := 2; i <= v/2; i++ {
			if v%i == 0 {
				flag += 1
				break
			} else {
				res = append(res, v)
			}
		}
	}
	return res
}

func main() {
	fmt.Println(PrimeNoSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}
