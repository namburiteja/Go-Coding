package main

import "fmt"

func main() {
	numbers := [5]int{10, 20, 30, 40, 50}

	fmt.Println("Length:", len(numbers))

	for i, value := range numbers {
		fmt.Printf("Index %d -> %d\n", i, value)
	}

	numbers[2] = 100

	fmt.Println("Updated Array:", numbers)
}