package main

import "fmt"

func main() {
	fmt.Println("Functions in go")

	greeter()
	greeterTwo()

	result := add(3, 5)
	fmt.Println("Result is", result)

	proResult, myMsg := proAdd(1, 2, 3, 4)
	fmt.Println("Pro result is : ", proResult)
	fmt.Println("Pro msg is : ", myMsg)
}

func greeterTwo() {
	fmt.Println("Another")
}

func greeter() {
	fmt.Println("Namaste from golang")
}

func add(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdd(values ...int) (int, string) {
	total := 0

	for _, v := range values {
		total += v
	}

	return total, "Hi pro add"
}
