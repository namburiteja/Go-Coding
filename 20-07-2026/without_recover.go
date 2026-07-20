package main
import (
	"fmt"
)
func processRequest() {
	panic("Database Connection Interrupted")
}
func main() {
	fmt.Println("server started")
	processRequest()
	fmt.Println("server still running")
}