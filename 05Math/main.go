package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	// "math/rand"
)

func main() {
	fmt.Println("Welcome to math in GOlang")

	// var myNumberOne int = 2
	// var myNumberTwo float64 = 4

	// fmt.Println("The sum is", myNumberOne + int(myNumberTwo))
	// TODO:INVALID

	// random number (math/rand)
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1) // between 0 and 4

	// crypto random

	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)

}
