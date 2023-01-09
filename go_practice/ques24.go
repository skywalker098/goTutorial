package main

import "fmt"

//Write a function that takes a slice of strings
//and returns a new slice with all strings that are at least a given length.
//Input: (["cat", "dog", "elephant", "lion"], 5) Output: ["elephant"]
func SliceWithLength(slice []string, n int) []string {
	res := []string{}
	for _, v := range slice {
		if len(v) >= n {
			res = append(res, v)
		}
	}
	return res
}

func main() {
	fmt.Println(SliceWithLength([]string{"cat", "dog", "elephant", "lion"}, 5))
}
