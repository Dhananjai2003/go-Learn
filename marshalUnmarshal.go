package main

import (
	"fmt"
	"encoding/json"
)

type Student struct{
	Roll int 
	Name string
}

func main() {

	student1 := Student{Roll:69 ,Name: "john Doe"}

	data, _ := json.Marshal(student1)
	fmt.Println(string(data))

	var StudentCopy Student

	json.Unmarshal(data,&StudentCopy)

	fmt.Println(StudentCopy.Name,StudentCopy.Roll)


}