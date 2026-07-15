package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30}

	fmt.Println("Original:", numbers)

	numbers = append(numbers, 40, 50)

	fmt.Println("After Append:", numbers)

	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))

	copySlice := make([]int, len(numbers))
	copy(copySlice, numbers)

	fmt.Println("Copied Slice:", copySlice)
}