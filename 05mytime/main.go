package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006"))
	fmt.Println(presentTime.Format("01-02-2006 Monday"))
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))


	createdDate := time.Date(2020, time.October, 31,20, 30, 0,0,time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
