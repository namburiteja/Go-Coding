package main
import (
	"fmt"
	"errors"
)
var val = errors.New("Invalid Password")
func login() error {
	return val
}
func authenticate() error {
	err := login()
	if err!=nil {
		return fmt.Errorf("Actual issue is : %w",err)
	}
	return nil
}
func main(){
	ans := authenticate()
	if errors.Is(ans,val) {
		fmt.Println("Password is wrong")
	}
}