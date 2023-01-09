package main

import "fmt"

///Write a function that takes a string and returns a new string
//with all the words in reverse order, with the letters in each word reversed as well.
//Input: "This is a test" Output: "tset a si sihT"
func StringReverseByLetter(str string) string {
	//reversing each character
	strLength := len(str)
	mid := strLength / 2
	var temp string = ""
	for i := 0; i < mid; i++ {
		temp = string(str[i])
		str[i] = str[strLength-i-1]
		str[strLength-i-1] = temp
	}
	return temp

}

func main() {
	res := 9 / 2
	fmt.Print(res)
}
