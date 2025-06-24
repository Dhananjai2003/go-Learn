package main

import (
	"fmt"
)

func counter() func() int {

	count := 0

	return func() int {
		count++
		return count
	}

}

func main() {

	countKeep := counter()

	fmt.Println(countKeep())
	fmt.Println(countKeep())
	fmt.Println(countKeep())


}