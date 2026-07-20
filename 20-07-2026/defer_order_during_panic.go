package main
import "fmt"
func A() {

	defer fmt.Println("A")

	B()

}

func B() {

	defer fmt.Println("B")

	panic("Boom")

}

func main() {

	A()

}