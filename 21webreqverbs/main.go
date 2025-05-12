package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to Web verb")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code", response.StatusCode)
	fmt.Println("Content length", response.ContentLength)

	var responseString strings.Builder

	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ :=responseString.Write(content)

	fmt.Println("ByteCount is", byteCount)
	fmt.Println(responseString.String())


	// fmt.Println(string(content))
}

func PerformPostJsonRequest(){
	const myurl = "http://localhost:8000/post"

	//fake json payload

	requestBody := strings.NewReader(`
		{
			"coursename":"Let's go with golang",
			"price":0,
			"platform":"Youtube"	
		}
	`)

	response,err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content,_:=ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFormRequest(){
	const myurl = "http://localhost:8000/postform"

	data := url.Values{}

	data.Add("firstname", "Sarthak")
	data.Add("lastname","Sharma")
	data.Add("email","sarthak@gmil.com")

	response,err := http.PostForm(myurl,data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
