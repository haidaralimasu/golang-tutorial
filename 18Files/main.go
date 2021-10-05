package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files handeling in go")

	content := "This needs to go in file - kodinghandle.com"

	file, err := os.Create("./myKHfile.txt")

	checkNillErr(err)

	length, err := io.WriteString(file, content)

	checkNillErr(err)

	fmt.Println("Length is ", length)

	readFile("./myKHfile.txt")

	defer file.Close()

}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)

	checkNillErr(err)

	fmt.Println("The data is ", string(databyte))
}

func checkNillErr(err error) {
	if err != nil {
		panic(err)
	}
}
