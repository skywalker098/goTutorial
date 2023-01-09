package main

import (
	"fmt"
	"strings"
)

// Write a function that takes a slice of strings
// and returns a new slice with all strings that contain a given letter.
// Input: (["cat", "dog", "elephant", "lion"], "e") Output: ["elephant"]
func SliceWithGivenLetter(slice []string, letter string) []string {
	res := []string{}
	for _, v := range slice {
		count := strings.Count(v, letter)
		if count >= 1 {
			res = append(res, v)
		}
	}
	return res
}

func main() {
	fmt.Println(SliceWithGivenLetter([]string{"cat", "dog", "elephant", "lion"}, "e"))
}
