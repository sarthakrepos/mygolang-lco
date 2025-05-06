package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])

	fmt.Println(fruitList)

	highScores := make([]int,4)

	highScores[0]=234
	highScores[1]=945
	highScores[2]=465 
	highScores[3]=867
	// highScores[4]=777
	highScores=append(highScores, 555,666,321)

	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

	var carList = []string{"Tata","Maruti","Mahindra","Toyota","Audi"}

	fmt.Println(carList)
	// carList = append(carList[:3],carList[4:]... )
	carList = append(carList[:2],append([]string {"Thar"}, carList[2:]...)... )
	fmt.Println(carList)

}
