package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")
	//no inheritance in golang
	//No super or parent

	sarthak :=User{"Sarthak","sarthak@gmail.com", true, 25}
	fmt.Println(sarthak)
	fmt.Printf("sarthak details are: %+v\n",sarthak)
	fmt.Printf("Name is %v, Email is %v\n ", sarthak.Name,sarthak.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
