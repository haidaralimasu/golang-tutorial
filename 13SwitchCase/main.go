package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to switch case")

	rand.Seed(time.Now().UnixNano())
	dice := rand.Intn(6) + 1

	fmt.Println("value of dice is ", dice)

	switch dice {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("Dice value is 2 spot")
	case 3:
		fmt.Println("Dice value is 3 spot")
		fallthrough
	case 4:
		fmt.Println("Dice value is 4 spot")
	case 5:
		fmt.Println("Dice value is 5 spot")
	case 6:
		fmt.Println("Dice value is 6 spot and a chance again")
	default:
		fmt.Println("what was that")
	}

}
