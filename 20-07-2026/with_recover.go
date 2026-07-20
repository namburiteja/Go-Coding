package main 
import (
	"fmt"
)
func processRequest() {
	defer func() {
		if r:=recover() ; r!=nil {
			fmt.Println("Recovered from panic :",r)
			fmt.Println("Returning internal server error")
		}
	}()
	panic("Database Connection Interrupted")
	fmt.Println("Request Completed")
}
func main() {
	fmt.Println("Server started")
	processRequest()
	fmt.Println("Server is still in running...")
}