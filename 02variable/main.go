package main

import "fmt"

func main()  {
	var username string
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n", username)


	var age int = 25
	fmt.Println(age)
	fmt.Printf("age is of type : %T \n", age)

	var isAdult bool = true
	fmt.Println(isAdult)
	fmt.Printf("isAdult is of type : %T \n", isAdult)
}