package main

import (
	"fmt"
	"strings"
)

// Write a function that takes a string and returns a new string with all vowels removed.
// Input: "hello" Output: "hll"
func NoVowelString(word string) string {
	var tempString string
	for _, v := range []string{"a", "A", "e", "E", "i", "I", "o", "O", "u", "U"} {
		tempString = strings.ReplaceAll(tempString, v, "")
	}
	return tempString
}

// func NoVowelString(word rune) string {
// 	var tempString rune
// 	for _, v := range word {
// 		if v != "A" || "E" || "I" || "O" || "U" || "a" || "e" || "i" || "o" || "u" {
// 			tempString += v
// 		}
// 	}
// 	return strconv.QuoteRune(tempString)
// }

func main() {
	// This won't work bcoz the input type should be rune
	// fmt.Println(NoVowelString("hello"))
	fmt.Println(NoVowelString("hello"))

}

/////////need to review
