package main

import "fmt"

var company = "Transcil" // Global variable

func greet() {
	name := "Teja" // Local variable
	fmt.Println("Name:", name)
	fmt.Println("Company:", company)
}

func main() {
	greet()
	fmt.Println("Company:", company)
}