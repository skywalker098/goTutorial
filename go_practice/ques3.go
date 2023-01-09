package main

import "fmt"

//Write a function that takes a string and returns a map with the frequency count of each character in the string.
//Input: "hello" Output: {"h": 1, "e": 1, "l": 2, "o": 1}
func FrequencyMap(word string) map[string]int {
	res := make(map[string]int)
	//chars := []rune(word)
	for _, j := range word {
		val, ok := res[string(j)]
		if ok {
			res[string(j)] = val + 1
		} else {
			res[string(j)] = 1
		}
	}
	return res
}

func main() {
	fmt.Println(FrequencyMap("hello"))
}
