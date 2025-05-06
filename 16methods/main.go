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
	sarthak.GetStatus()
	sarthak.NewEmail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus(){
	fmt.Println("Is user active", u.Status)
}

func (u User) NewEmail()  {
	u.Email="s@gmail.com"
	fmt.Println("New Mail is : ", u.Email)
}
