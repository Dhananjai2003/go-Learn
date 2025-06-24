package main

import (
	"fmt"
)

func cal(a int, b int) (int, int) {

	return a+b , a-b

}

func main() {

	a , b := cal(2,1)

	fmt.Println(a,b)

}