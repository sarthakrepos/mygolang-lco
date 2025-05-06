package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golangs")
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[3] = "Papaya"
	fmt.Println("Fruit List is: ", fruitList)

	var vegList = [3]string{"potato","beans","mushroom"}
	fmt.Println(len(vegList))
}
