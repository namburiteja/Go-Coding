package main

import "fmt"

func main() {

	// Anonymous function assigned to a variable
	add := func(a, b int) int {
		return a + b
	}

	fmt.Println("Addition:", add(10, 20))

	// Immediately Invoked Function Expression (IIFE)
	func() {
		fmt.Println("Hello from Anonymous Function!")
	}()
}