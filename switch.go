package main

import (

	"fmt"

)

func main() {

	i := 2
	switch i {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	default:
		fmt.Println("Not 1 Not 2")
		return
	}

	i = 4

	switch i {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	default:
		fmt.Println("Not 1 Not 2")
		return
	}



}