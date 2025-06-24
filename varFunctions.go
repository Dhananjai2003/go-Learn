package main

import (
	"fmt"
)

func sum(n ...int) int { // as many variables as required 
	total := 0

	for _, num := range n {
		total += num		
	}

	return total
}

func main() {

	fmt.Println(sum(1,2,3,4))
	fmt.Println(sum(1,2,3))

}