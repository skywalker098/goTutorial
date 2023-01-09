package main

import "fmt"

//Pointer is the refernce to the direct memory location
func main() {
	fmt.Println("Welcome to pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr)

	myNumber := 22

	var ptr = &myNumber
	fmt.Println("Value of pointer is ", ptr)  //address of the pointer
	fmt.Println("Value of pointer is ", *ptr) //value inside the pointer

	*ptr = *ptr + 2

	fmt.Println("New value is: ", myNumber)
}
