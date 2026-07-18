package main
import (
	"fmt"
	"os"
)
func main() {
	err := os.Rename("assignments","assignments_renamed")
	if err!=nil {
		fmt.Println(err)
	}else {
		fmt.Println("Directory renamed successfully")
	}
}