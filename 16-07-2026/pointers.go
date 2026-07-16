package main

import "fmt"

func main() {
	a := 10
	p := &a
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(p)
	fmt.Println(*p)

}