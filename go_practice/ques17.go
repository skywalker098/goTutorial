package main

import "fmt"

//Write a function that takes a slice of strings and returns a new slice with all strings that start with a given letter.
//Input: (["cat", "dog", "elephant"], "e") Output: ["elephant"]

func StringWithLetter(slice []string, str string) []string {
	tempSlice := []string{}
	for _, v := range slice {
		for i := 0; i < len(slice); i++ {
			m := string(v[i])
			if m == str {
				tempSlice = append(tempSlice, v)
				break
			}
		}
	}
	return tempSlice
}

func main() {
	fmt.Println(StringWithLetter([]string{"cat", "dog", "elephant"}, "e"))
}
