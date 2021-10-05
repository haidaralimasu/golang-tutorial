package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Handeling web request in go lang")

	response, err := http.Get(url)

	checkNillErr(err)

	fmt.Printf("response is of type %T\n ", response)

	defer response.Body.Close() // close the connection

	data, err := ioutil.ReadAll(response.Body)

	checkNillErr(err)

	content := string(data)
	fmt.Println("Response is \n", content)

}

func checkNillErr(err error) {
	if err != nil {
		panic(err)
	}
}
