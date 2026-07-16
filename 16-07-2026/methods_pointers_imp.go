package main
import "fmt"
type Rectangle struct {
	length int
	width int
}
func (x Rectangle) area() int {
	return x.length * x.width
}
func (x Rectangle) perimeter() int {
	return 2 * (x.width + x.length)
}
func (x *Rectangle) changeDimensions(length int, width int) {
	x.length = length
	x.width = width
	fmt.Println("Inside changeDimensions method : ",x.length,x.width)
}
func main() {
	r := Rectangle{
		length : 10,
		width : 20,
	}
	fmt.Println("Area of Recatngele is : ",r.area())
	fmt.Println("Perimeter of Rectangle is : ",r.perimeter())
	r.changeDimensions(30,40)
	fmt.Println("Area of Recatngele is : ",r.area())
	fmt.Println("Perimeter of Rectangle is : ",r.perimeter())	
}
