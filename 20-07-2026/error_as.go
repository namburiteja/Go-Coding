package main
import (
	"fmt"
	"errors"
)
type validator struct {
	field string
}
func (v validator) Error() string {
	return fmt.Sprintf("%s is wrong",v.field)
}
func register() error {
	return validator{
		field : "Email",
	} 
}
func service() error{
	err := register()
	if err!=nil {
		return fmt.Errorf("registration failed : %w",err)
	}
	return nil
}
func main() {
	err := service()
	var val validator
	if errors.As(err,&val){
		fmt.Println(val.field)
	}
}