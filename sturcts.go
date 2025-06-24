package main

import (
	"fmt"
)

type student struct {
	name string
	age int 
}

func addStudent(name string, age int) *student {
	p := student{name: name, age: age}
	return &p
}

func main() {

	fmt.Println(addStudent("John",13))

	student1 := addStudent("snow",48)
	fmt.Println(student1)

	student2 := student{name : "bob", age: 10}
	fmt.Println(student2)

	newDog := struct{
		name string
		age int 
	}{
		"ruby",
		3,
	}

	fmt.Println(newDog)

}