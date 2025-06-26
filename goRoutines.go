package main

import (
	"fmt"
	"time"
)

func task(id int) {

	for i := 0; i<5 ; i++ {
		fmt.Printf("This line is from task id : %d iteration number : %d\n",id,i+1)
		time.Sleep(500*time.Millisecond)
	}

}

func main() {
	for i:=0; i<5; i++{
		go task(i)
	}

	time.Sleep(2*time.Second)
}