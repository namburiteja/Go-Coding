package main

import "fmt"

func main() {
	marks := make(map[string]int)

	marks["Teja"] = 90
	marks["Rahul"] = 85
	marks["Ravi"] = 95

	fmt.Println("Marks:", marks)

	marks["Teja"] = 100

	delete(marks, "Rahul")

	if mark, exists := marks["Teja"]; exists {
		fmt.Println("Teja's Marks:", mark)
	}

	for name, mark := range marks {
		fmt.Println(name, "->", mark)
	}
}