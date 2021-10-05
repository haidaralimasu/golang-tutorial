package main

import "fmt"

// jwtToken := 3000000 (not allowed outside with valouros opertaotr :=)

const LoginToken string = "hiialdbafseufdi" // public

func main() {
	var userName string = "haidar"
	fmt.Println(userName)
	fmt.Printf("Variable is of type: %T \n", userName)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.455545435849358 //5 values
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var smallFloat64 = 255.9034903
	fmt.Println(smallFloat64)
	fmt.Printf("Variable is of type: %T \n", smallFloat64)

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	var anotherString string
	fmt.Println(anotherString)
	fmt.Printf("Variable is of type: %T \n", anotherString)

	var anotherBool bool
	fmt.Println(anotherBool)
	fmt.Printf("Variable is of type: %T \n", anotherBool)

	// implicit type

	var website = "kodinghandle.com"
	fmt.Println(website) // lexer ke jalwe

	numberOfUser := 300000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
