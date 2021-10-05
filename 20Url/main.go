package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=afudifid5"

func main() {
	fmt.Println("handeling url in go-lang")

	fmt.Println(myurl)

	//parsing
	result, err := url.Parse(myurl)

	if err != nil {
		panic(err)
	}

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Printf("%T\n", qparams)
	fmt.Println(qparams)

	for _, v := range qparams {
		fmt.Println("pram is ", v)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=haidar",
	}

	anotherUrl := partsOfUrl.String()

	fmt.Println(anotherUrl)
}
