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
	// EncodeJson()
	DecodeJson()
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

func DecodeJson() {
	jsonDataFormWeb := []byte(`
	
	{
		"coursename": "ReactJs Bootcamp",
		"courseprice": 299,
		"website": "LearnCodeOnline.in",
		"tags": [
				"web-dev",
				"js"
		]
	}
	
`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFormWeb)

	if checkValid {
		fmt.Println("valid")
		json.Unmarshal(jsonDataFormWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("not valid")
	}

	// some cases where you just want key value pair

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFormWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is : %T \n", k, v, v)
	}

}
