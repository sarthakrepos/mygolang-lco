package main

import "fmt"

func main() {
	fmt.Println("Function in go")
	greeter()

	// result := adder(3,5)

	fmt.Println("Result is: ", adder(3,5))

	proResult := proAdder(2,5,7,9,4,1,7)

	fmt.Println("proResut is: ", proResult)

	resultSub,mymsg := sub(5,3)

	fmt.Println("Result is: ", resultSub)
	fmt.Println("Result is: ", mymsg)

	
	greeterTwo()
}

func adder(valOne int, valTwo int) int {
	return valOne+valTwo
}

func sub(valOne int, valTwo int) (int,string) {
	return valOne-valTwo,"Hello ji"
}

func proAdder(values...int) int {
	total := 0

	for _,val := range values{
		total += val
	}
	return total
}


func greeter(){
	fmt.Println("Namastey from golang")
}

func greeterTwo()  {
	fmt.Println("Another function")
}
