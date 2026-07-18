package main
import (
	"fmt"
	"errors"
)

func main() {
	err := errors.New("This is an custom error")
	fmt.Println(err)
}