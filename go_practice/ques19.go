package main

import (
	"fmt"
	"strings"
)

//Write a function that takes a string and returns a new string with all the words in reverse order.
//Input: "This is a test" Output: "test a is This"

func StringReversed(str string) string {
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

	tempString := strings.Split(str, " ")
	res := []string{}
	for i := len(tempString) - 1; i >= 0; i-- {
		res = append(res, tempString[i])
	}
	return strings.Join(res, " ")

}

func main() {
	fmt.Println(StringReversed("This is a test"))
}
