package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang")

	myCh := make(chan int,2)
	wg := &sync.WaitGroup{}

	// myCh <- 5
	// val:=<-myCh
	// fmt.Println(val)
	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-myCh)
		// fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		myCh <-5
		myCh <-6
		close(myCh)
		wg.Done()
	}(myCh, wg)

	wg.Wait()

}
