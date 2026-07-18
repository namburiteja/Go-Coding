package main 
import (
	"fmt"
	"os"
)
func main() {
	err := os.Mkdir("assignments",0755)
	if err!= nil {
		fmt.Println(err)
	}else {
		fmt.Println("Directory created successfully")
	}
}