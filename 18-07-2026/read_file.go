package main
import (
	"fmt"
	"os"
) 
func main() {
	data,err := os.ReadFile("read.txt")
	if err!=nil {
		fmt.Println(err)
	}else{
		fmt.Println(string(data))
	
	}
}