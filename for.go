package main

import (
	"fmt"
)

func main() {

	i := 1 
	for i <= 5{
		for j := 1; j <= i; j++{
			fmt.Print('x')
		}
		fmt.Println()
		i++
	} 

	for k := 0; ; k++ {
		if k % 2 == 1 {
			continue
		}

		fmt.Println(k)

		if k == 6 {
			break
		}

	}
//learned if and for loop conceptss

}