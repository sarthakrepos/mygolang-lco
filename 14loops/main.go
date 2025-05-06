package main

import "fmt"

func main() {
	fmt.Println("Welcome to loop")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	for day := 0; day < len(days); day++ {
		fmt.Println(days[day])
	}

	//this will also return index
	for i := range days {
		fmt.Println(days[i])
	}

	// for each type will return value
	for index, day := range days {
		fmt.Printf("index is %v and value us %v\n", index, day)
	}

	for _, day := range days {
		fmt.Printf("value us %v\n", day)
	}

	rougeValue := 1

	for rougeValue < 10 {
		if rougeValue ==2 {
			goto lco
		}
		if rougeValue == 5 {
			break
		}
		fmt.Println(rougeValue)
		rougeValue++

	}
lco:
	fmt.Println("Welcome")
}
