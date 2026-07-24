package main
import (
	"fmt"
	"time"
)
func hello() {
	fmt.Println("This is Go routine")
}
func main() {
	fmt.Println("Hello START")
	go hello()
	time.Sleep(time.Second)
	fmt.Println("After Go routine")
}