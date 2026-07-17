package main
import "fmt"

type Vehicle interface {
	drive()
}
type car struct {}
type bike struct {}
func (c car) drive() {
	fmt.Println("car is Driving")
}
func (b bike) drive() {
	fmt.Println("bike is Driving")
}
func driving(v Vehicle) {
	v.drive()
}
func main() {
	c := car{}
	b := bike{}

	driving(c)
	driving(b)
}