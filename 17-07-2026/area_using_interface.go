package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	length, width float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func printArea(s Shape) {
	fmt.Println("Area:", s.Area())
}

func main() {
	r := Rectangle{
		length: 10,
		width:  3,
	}

	c := Circle{
		radius: 5,
	}

	fmt.Println("Area of rectangle:", r.Area())
	fmt.Println("Area of circle:", c.Area())
}