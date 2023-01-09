package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//Generating a random number and storing in variable "randNumber"
	rand.Seed(time.Now().UnixNano())
	randNumber := rand.Intn(100)
	//Printing the random Number
	fmt.Print("Generated Radndom Number: ", randNumber)

	//Taking input from user
	attempt := (10 - 1)

	for attempt >= 0 {

		fmt.Print("\nEnter the Number: ")
		var num int
		fmt.Scanln(&num)

		//Matching the input with genrated random number
		if num > randNumber {
			fmt.Printf("Number is too high, Try again (attempt left: %d)", attempt)
		} else if num < randNumber {
			fmt.Printf("Number is too low, Try again (attempt left: %d)", attempt)
		} else {
			fmt.Printf("You guessed the number!!")
		}
		attempt--
	}

}
