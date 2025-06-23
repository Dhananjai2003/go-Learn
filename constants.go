package main

import (
	"fmt"
	"math"
)

const s string = "Hello i am a constant"

func main() {

	fmt.Println(s) // prints "Hello i am a constant"

	const n = 500000

	const d = 3e20/n
	fmt.Println(d)
	fmt.Println(int64(d)) // type cast to 64 bit ints

	fmt.Println(math.Sin(n)) //here the sin function expects float64		

}