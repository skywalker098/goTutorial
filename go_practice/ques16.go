package main

import "fmt"

//Write a function that takes a slice of strings and a string,
//and returns the index of the first occurrence of the string in the slice.
//Input: (["cat", "dog", "elephant"], "dog") Output: 1
func FirstOccurenceOfString(slice []string, str string) int {

	for i := 0; i < len(slice); i++ {
		if slice[i] == str {
			return i
			break
		}
	}
	return -1
}

func main() {
	fmt.Println(FirstOccurenceOfString([]string{"cat", "dog", "elephant"}, "dog"))
}
