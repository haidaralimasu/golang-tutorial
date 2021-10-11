package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"courseprice"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("go lang json")
	EncodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJs Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "asdc123", []string{"fullstack", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "cdbc123", nil},
	}
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	// fmt.Println(finalJson)
	fmt.Printf("%s\n", finalJson)
}
