package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops in golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wedensday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for _, day := range days {
		fmt.Println(day)
	}

	rougeValue := 1

	for rougeValue < 10 {

		if rougeValue == 2 {
			goto kh
		}

		fmt.Println("Valur is : ", rougeValue)
		rougeValue++
	}

kh:
	fmt.Println("kodinghandle.com")
}
