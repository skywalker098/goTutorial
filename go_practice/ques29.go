package main

import "fmt"

//Write a function that takes a slice of integers and
//returns a new slice with all elements that are palindromes (i.e., read the same forwards as backwards).
//Input: [1, 11, 121, 12321] Output: [11, 121, 12321]

//Helper function to individually check slice element for palindrome
func isNumPalindrome(num int) bool {
	var remainderNum int
	oriNum := num
	res := 0
	if num < 9 {
		return false
	}
	for num > 0 {
		remainderNum = num % 10
		res = (res * 10) + remainderNum
		num = num / 10
	}
	if oriNum == res {
		return true
	} else {
		return false
	}
}

func CheckPalindromeNum(slice []int) []int {
	tempSlice := []int{}
	for _, value := range slice {
		if isNumPalindrome(value) == true {
			tempSlice = append(tempSlice, value)
		}
	}
	return tempSlice
}

func main() {
	fmt.Println(CheckPalindromeNum([]int{1, 11, 121, 12321}))
}

// func main() {
// 	fmt.Println(isNumPalindrome(1))
// }
