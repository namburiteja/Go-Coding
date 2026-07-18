package main 
import "fmt"

type Animal interface {
	type a int = 10
	sleep()
}
type Dog struct {}
func (d *Dog) sleep() {
	fmt.Println("Dog is sleeping")
}
func main() {
	d := Dog{}
	d.sleep()
	var a Animal = &d
	a.sleep()

}