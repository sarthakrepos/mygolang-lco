package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to Json video")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourse := []course{
		{"ReactJs", 299, "Youtube", "abc123", []string{"web-dev", "js"}},
		{"Mern", 199, "Youtube", "abc12", []string{"full-stack", "js"}},
		{"Angular", 99, "Youtube", "abc1", nil},
	}

	//package this data as Json data

	// finalJson,err := json.Marshal(lcoCourse)
	finalJson, err := json.MarshalIndent(lcoCourse, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	
		{
                "coursename": "ReactJs",
                "Price": 299,
                "Platform": "Youtube",
                "tags": [
                        "web-dev",
                        "js"
                ]
        }
	`)

	var lcoCourse course
	chackValid := json.Valid(jsonDataFromWeb)

	if chackValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("Json was not valid")
	}

	// some cases where you just ant to add data to key value

	var myOnlineData map[string]interface{}

	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k,v := range myOnlineData{
		fmt.Printf("Key is %v and value is %v and Type is : %T\n"	,k,v,v)
	}
}
