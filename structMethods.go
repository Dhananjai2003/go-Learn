package main

import (
	"fmt"
)

type student struct {
	name string
	age int
	marks int
}

func (s *student) passOrNot() bool {
	if s.marks > 40 {
		return true
	}

	return false
}

func main() {

	student1 := student{name: "john", age: 18, marks: 41}
	fmt.Println(student1.passOrNot())

	student2 := student{name: "doe", age: 24, marks:20}
	fmt.Println(student2.passOrNot())


}