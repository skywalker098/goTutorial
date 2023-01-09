package main

// Pointer method without return
func multiply(number *int) {
	*number *= 2
}

//Pointer method with return
func sum(number *int) *int {
	value := *number + 10
	return &value
}

// func main() {
// 	number := 10
// 	multiply(&number)
// 	fmt.Println(number)

// 	s := sum(&number)
// 	fmt.Println(*s)
// }
