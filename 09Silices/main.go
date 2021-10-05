package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	var fruitsList = []string{"apple", "tomato", "peach"}
	// when added value in [] is array and if [] is empty it is slice

	fmt.Printf("The type of fruitList is %T\n", fruitsList)

	fruitsList = append(fruitsList, "Mango", "Banana")
	fmt.Println(fruitsList)

	fruitsList = append(fruitsList[:3]) // first element will removed
	fmt.Println(fruitsList)

	highScore := make([]int, 4)

	highScore[0] = 234
	highScore[1] = 945
	highScore[2] = 465
	highScore[3] = 867
	fmt.Println(sort.IntsAreSorted(highScore))

	highScore = append(highScore, 555, 666, 321) // append re allocate memory

	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println(highScore)

	fmt.Println(sort.IntsAreSorted(highScore))

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
