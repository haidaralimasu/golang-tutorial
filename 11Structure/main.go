package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to struct in golang")
	// no inheritance in golang: No super or parent

	haidar := User{"Haidar", "haidar@go.dev", true, 19}
	fmt.Println(haidar)
	fmt.Printf("haidar details are: %+v\n", haidar)
	fmt.Printf("Name is %v and email is %v\n", haidar.Name, haidar.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
