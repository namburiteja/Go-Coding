package main
import (
	"fmt"
	"strconv"
)
func main() {
	n,err := strconv.Atoi("abc")
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}