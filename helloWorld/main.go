package main

import (
	"fmt"
	"math/rand"
	"time"
)
func.main() {
	//fmt.Println("Hello, Workd")
	//Genrerate a random number between 1 and 100
	rand.seed(time.Now().Unix())
	seccretNumber := rand.Intn(100)



	//read the input from sdin (terminal)
	fmt.Print("Enter a number: ")
	var input int
	
	_, err := fmt.scan(&input)
	if err != nil {
		fmt.Println("Error in reading  input")
		return
	}

	//Print the input
	fmt.Println(input)

	



}