package main

import (
	"fmt"
	"strings"
)

//Write a function that takes a string and returns a new string with all the words in reverse order.
//Input: "This is a test" Output: "test a is This"

// func StringReversed(str string) string {
//reversing the string using temp variable to swap the n and len(string)-n till we reach the mid
//reversing each character
// strLength := len(str)
// mid := strLength / 2
// var temp string = ""
// for i := 0; i < mid; i++ {
// 	temp = string(str[i])
// 	str[i] = str[strLength-i-1]
// 	str[strLength-i-1] = temp
// }
// return temp

func PrintBackwards(s string) string {
	part := strings.Split(s, " ")
	res := []string{}
	for i := len(part) - 1; i >= 0; i-- {
		if part[i] == " " {
			continue
		}
		res = append(res, part[i])
	}
	return strings.Join(res, " ")
}

func main() {
	fmt.Println(StringReversed("This is a test"))
}
