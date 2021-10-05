package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()
}

// last in first out

// defer execute in reverse

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
