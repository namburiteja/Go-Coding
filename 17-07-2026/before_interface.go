package main
import "fmt"

type Dog struct{}

type Cat struct{}

func (d Dog) Speak() {
	fmt.Println("Woof")
}
func (c Cat) Speak() {
	fmt.Println("Meow")
}
func MakeDogSpeak(d Dog) {
	d.Speak()
}
func MakeCatSpeak(c Cat) {
	c.Speak()
}
func main() {
	dog := Dog{}
	cat := Cat{}

	MakeDogSpeak(dog)
	MakeCatSpeak(cat)
}