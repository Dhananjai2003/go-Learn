package main

import (
	"fmt"
)

func main() {

	a := []int{1,2,3}
	sum := 0

	for _,i := range a{
		sum+=i
	}

	fmt.Println(sum)

	dict := map[string]string{"a":"Dhananjai","b":"John"}

	for k,v := range dict{
		fmt.Printf("%s : %s\n",k,v)
	}

}