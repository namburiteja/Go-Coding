package main
import "fmt"

type MyInt int
 
func (m MyInt) isEven() bool {
	if m%2==0 {
		return true
	}
	return false
}
func main() {
	m := MyInt(10)
	fmt.Println("Is Even :",m.isEven())

}