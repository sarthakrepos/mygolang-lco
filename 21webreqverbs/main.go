package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Welcome to Web verb")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "https://jsonplaceholder.typicode.com/todos/1"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code", response.StatusCode)
	fmt.Println("Content length", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
