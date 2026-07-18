package main
import "fmt"

type Sleep interface {
	sleep()
}
type Eat interface {
	eat()
}
type Animal interface {
	Sleep
	Eat
}

type Dog struct {}
func (d Dog) sleep() {
	fmt.Println("Dog is Sleeping")
}
func (d Dog) eat() {
	fmt.Println("Dog is Eating")
}
func main() {
	d := Dog{}
	d.sleep()
	d.eat()
}