package main

import "fmt"

//Write a function that takes a slice of strings and a string,
//and returns true if the string is contained in the slice.
//Input: (["cat", "dog", "elephant"], "dog") Output: true
func StringMatch(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(StringMatch([]string{"cat", "dog", "elephant"}, "dog"))
}
