package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("GET request in golang")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl string = "http://localhost:3000"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("status : ", response.StatusCode)
	fmt.Println("length ", response.ContentLength)

	var responseString strings.Builder
	content, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	byteCount, _ := responseString.Write(content)

	fmt.Println(byteCount)
	fmt.Println(responseString.String())




	// fmt.Println(string(content))
	// fmt.Println(content)
}
