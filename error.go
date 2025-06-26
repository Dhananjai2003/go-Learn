package main

import (
	"fmt"
	"errors"
)


func divide(a float32, b float32) (float32,error){
	if b == 0 {
		return -1, errors.New("Cannot divide by 0")
	}

	return a / b, nil
}

func main() {

	a, b := divide(10,2)
	fmt.Println(a,b)

	a,b = divide(10,0)
	fmt.Println(a,b)

}
