package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// fmt.Println("Hello, World1!")
	rand.Seed(time.Now().UnixNano())
	randNumber := rand.Intn(100)
	fmt.Println(randNumber)
}
