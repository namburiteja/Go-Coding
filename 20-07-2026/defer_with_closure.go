package main
import "fmt"
func main() {
	x:=10
	defer func() {
		fmt.Println(x)
	}()
	x = 20
}