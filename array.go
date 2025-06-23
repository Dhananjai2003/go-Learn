package main

import (
	"fmt"
)

func main() {

	var a [5]int
	fmt.Println(a)

	a[4] = 10	
	fmt.Println(a)

	fmt.Println(len(a))

	var TwoD [5][5]int

	
	for i := 0; i < 5; i++{
		for j := 0; j < 5; j++ {
			TwoD[i][j]=i+j;
		}
	}

	fmt.Println(TwoD)

}