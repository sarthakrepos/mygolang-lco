package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")

	content := "This needs to go in a file"

	file , err := os.Create("./mygofile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Length is ", length)

	defer file.Close()

	readFile("./mygofile.txt")

}

func readFile(filename string)  {
	databyte, err := ioutil.ReadFile(filename)
	// if err != nil {
	// 	panic(err)
	// }
	checkError(err)

	fmt.Println("Text data inside file is \n", string(databyte))
}

func checkError(err error){
	if err != nil {
		panic(err)
	}
}