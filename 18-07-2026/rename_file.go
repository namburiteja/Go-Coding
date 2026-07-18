package main
import (
	"fmt"
	"os"
)
func main() {
	err := os.Rename("create_new_directory.go","create_directory.go")
	if err!=nil {
		fmt.Println(err)
	}else { 
		fmt.Println("File name Renamed Successfully")
	}
}