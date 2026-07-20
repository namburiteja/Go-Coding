package main
import (
	"fmt"
	"errors"
)
func check() error {
	return errors.New("Database Error")
}
func service() error {
	err := check()
	if err!=nil {
		return fmt.Errorf("The Actual Error is : %w",err)
	}
	return  nil
}
func main() {
	err := service()
	if err!=nil {
		fmt.Println(err)
	}
}