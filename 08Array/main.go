package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array in go lang")

	var fruitsList [4]string

	fruitsList[0] = "Apple"
	fruitsList[1] = "Tomato"
	fruitsList[3] = "Peach"

	fmt.Println("Fruit list is ", fruitsList)
	fmt.Println("Fruit list length is ", len(fruitsList))

	var vegList = [3]string{"potato", "beans", "mushroom"}

	fmt.Println("veggie list is ", vegList)
	fmt.Println("veggie list length is", len(vegList))
}
