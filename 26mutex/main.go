package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup

func main() {

	websitelist := []string{
		"https://google.com",
		"https://facebook.com",
		"https://chatgpt.com",
		"https://github.com",
		"https://apple.com",
	}
	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in endpoint")
	}else{
		signals = append(signals, endpoint)
		fmt.Printf("%d 200 status code for website %s\n", res.StatusCode, endpoint)
	}

	
}
