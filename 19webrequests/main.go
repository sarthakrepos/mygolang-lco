package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://google.com"

func main() {
	fmt.Println("Web requests")

	response,err :=http.Get(url)
	if err!=nil{
		panic(err)
	}
	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close() //caller's responsibility to close the connection

	dataBytes,err:=ioutil.ReadAll(response.Body)
	if err!=nil {
		panic(err)
	}

	content :=string(dataBytes)

	fmt.Println(content)

}
