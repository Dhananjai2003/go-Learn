package main

import (
	"fmt"
)

type shape interface{
	area() float32
}

type rectange struct {
	lenght float32
	width float32
}

func (rect rectange) area() float32 {
	return rect.lenght * rect.width
}

type circle struct {
	radius float32
}

func (cir circle) area() float32 
{
	return 3.14 * cir.radius * cir.radius
}

func printArea(s shape){
	fmt.Println(s.area())
}

func main() {


	rect1 := rectange{lenght: float32(10), width: float32(12)}
	cir1 := circle{radius: float32(10)}

	printArea(rect1)
	printArea(cir1)


}