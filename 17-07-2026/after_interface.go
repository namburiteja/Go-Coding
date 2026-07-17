package main

import "fmt"

type Animal interface {
	Speak()
}

type Dog struct{}

func (d Dog) Speak() {
	fmt.Println("Woof")
}

type Cat struct{}

func (c Cat) Speak() {
	fmt.Println("Meow")
}

func MakeSpeak(a Animal) {
	a.Speak()
}

func main() {

	dog := Dog{}
	cat := Cat{}

	MakeSpeak(dog)
	MakeSpeak(cat)

}