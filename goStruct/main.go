package main

import (
	"fmt"
)

//structures in go

type User struct {
	Name   string
	Age    int
	Email  string
	Status bool
}

func main() {
	fmt.Println("--struct in go--")
	rohan := User{"Rohan", 23, "rohan@go.com", true}

	//use "Printf" inplace of "Println" when formatting with "%v" or "%+v"
	fmt.Printf("Details are: %+v\n", rohan)
	fmt.Printf("Name is %v and email is %v", rohan.Name, rohan.Email)
}
