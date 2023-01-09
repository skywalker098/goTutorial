package main

import "fmt"

//Write a function that takes a slice of strings and returns a new slice with all strings that have more than 5 characters.
//Input: ["cat", "dog", "elephant", "lion"] Output: ["elephant"]
func BigSLice(slice []string) []string {
	tempSLice := []string{}
	for i := 0; i < len(slice); i++ {
		if len(slice[i]) >= 5 {
			tempSLice = append(tempSLice, slice[i])
		}
	}
	return tempSLice
}

func main() {
	fmt.Println(BigSLice([]string{"cat", "dog", "elephant", "lion"}))
}
