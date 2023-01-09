package main

import (
	"fmt"
	"sort"
)

// Write a function that takes a slice of strings and returns a new slice with all strings sorted in alphabetical order.
// Input: ["zebra", "monkey", "aardvark"] Output: ["aardvark", "monkey", "zebra"]
func SortedSlice(slice []string) []string {
	sort.Strings(slice)
	return slice
}

func main() {
	fmt.Println(SortedSlice([]string{"zebra", "monkey", "aardvark"}))
}
