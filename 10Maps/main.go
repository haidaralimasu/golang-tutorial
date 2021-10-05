package main

import "fmt"

func main() {
	fmt.Println("Maps in go lang")

	// creating maps
	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("JS shorts for ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all languages: ", languages)

	// loops are interseting in go
	// like for each
	for _, value := range languages {
		fmt.Printf("For key v value is %v\n", value)
	}

}
