package main

import (
	"fmt"
	"os"
)



func main() {

	text := "Hello this is a sample text to be stored in a file through go"


	err := os.WriteFile("sample.txt", []byte(text), 0644) // 0644 is permission for standard read/write

	if err!=nil {
		fmt.Println("Some error ouccured during write")
		return
	}

	content, err := os.ReadFile("sample.txt") //reads in bytes

	if err!= nil{
		fmt.Println("Some error ouccured during read")
		return
	}

	fmt.Println(string(content))



}