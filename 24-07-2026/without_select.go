package main
import (
	"fmt"
	"time"
)
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		time.Sleep(8 * time.Second)
		ch1 <- "Worker 1"
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- "Worker 2"
	}()
	fmt.Println(<-ch1)
	fmt.Println(<-ch2)
}