package main

import (
	"fmt"
)

type ArithematicError struct {
	code int 
	description string
}

func (err ArithematicError) Error() string {

	return fmt.Sprintf("Error : %d %s",err.code,err.description)

}

func divide(a float32, b float32) (float32, error) {

	if b == 0 {
		return -1, ArithematicError{code: 400, description: "cannot divide by ZERO"}
	}

	return a / b, nil

}


func main() {

	ans , err := divide(10,0)

	if err != nil {

		fmt.Println(err)

		if myErr , ok := err.(ArithematicError); ok {
			fmt.Println("Error Code : ",myErr.code);
		}

	} else{
		fmt.Println("result is ",ans)
	}

}