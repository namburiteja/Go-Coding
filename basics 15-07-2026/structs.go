package main

import "fmt"

type Student struct {
	Name   string
	Age    int
	Branch string
	CGPA   float64
}

func main() {
	student := Student{
		Name:   "Teja",
		Age:    22,
		Branch: "CSE",
		CGPA:   8.8,
	}

	fmt.Println(student)

	fmt.Println("Name:", student.Name)
	fmt.Println("Age:", student.Age)
}