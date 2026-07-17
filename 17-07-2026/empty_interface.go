package main
import "fmt"
var value interface{}
func main() {
	value = 10
	fmt.Println("value is of type int:", value)

	value = "Hello"
	fmt.Println("value is of type string:", value)	

	value = 3.14
	fmt.Println("value is of type float64:", value)

	value = true
	fmt.Println("value is of type bool:", value)

	value = []int{1, 2, 3}
	fmt.Println("value is of type slice of int:", value)
}