package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// go greeter("Hello")
	// greeter("world")
	websitelist := []string{
		"https://google.com",
		"https://facebook.com",
		"https://go.dev",
		"https://github.com",
		"https://apple.com",
	}
	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in endpoint")
	}

	fmt.Printf("%d 200 status code for website %s\n", res.StatusCode, endpoint)
}
