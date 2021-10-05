package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers in golang")

	// var ptr *int // declaring pointer

	// fmt.Println("Value of pointer is", ptr)

	myNumber := 23

	var ptr = &myNumber // pointer refrencing means &

	fmt.Println("Value of actual pointer is ", ptr)  // memory location
	fmt.Println("Value of actual pointer is ", *ptr) // value of pointer using *

	*ptr = *ptr + 2
	fmt.Println("New value is ", myNumber)

}
