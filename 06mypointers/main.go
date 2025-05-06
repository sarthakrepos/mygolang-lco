package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointer")

	var ptr *int
	fmt.Println("Value of pounter is ", ptr)


	myNumber := 23
	var pt = &myNumber
	fmt.Println("Address is: ", pt)
	fmt.Println("Value is: ", *pt)

	*pt=*pt+2
	fmt.Println("Address is: ", pt)
	fmt.Println("Value is: ", *pt)

}
