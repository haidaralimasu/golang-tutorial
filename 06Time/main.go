package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime.Format("01 - 02 - 2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.August, 10, 28, 23, 04, 04, time.UTC)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))
}
