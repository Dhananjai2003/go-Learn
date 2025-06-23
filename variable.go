package main

import (
	"fmt"
)

func main(){
	var firstName = "John"
	var secondName = "Doe"

	fmt.Println(firstName + secondName) //prints "JohnDoe"

	var a, b int = 1,2
	fmt.Println(a,b) //prints "1 2"

	var j int
	fmt.Println(j) // prints '0'

	f := "hello"
	fmt.Println(f) // prints "hello"
}