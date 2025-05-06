package main

import "fmt"

func main() {
	fmt.Println("If else in go")
	
	var output string
	value:=9

	if value>10{
		output ="greater than 10"
	}else{
		output="less than 10"
	}

	fmt.Println(output)

}
