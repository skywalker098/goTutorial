package main

import "fmt"

//Write a function that takes a slice of strings and returns a new slice with all strings that are palindromes (i.e., read the same backwards as forwards).
//Input: ["racecar", "level", "hello"] Output: ["racecar", "level"]
func isPalindrome(str string) bool {
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	for i := range str {
		if str[i] != reversedStr[i] {
			return false
		}
	}
	return true
}

func PalindromeStringCheck(slice []string) []string {
	tempSlice := []string{}
	for _, value := range slice {
		if isPalindrome(value) == true {
			tempSlice = append(tempSlice, value)
		}
	}
	return tempSlice
}
func main() {
	fmt.Println(PalindromeStringCheck([]string{"racecar", "level", "hello"}))
}
