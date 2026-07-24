package main
import (
	"fmt"
	"time"
)
func work(i int) {
		fmt.Println("worker",i*i)
}
func main() {
	fmt.Println("Now started")
	for i:=0;i<10;i++ {
		go work(i)
		time.Sleep(3*time.Second)

	}
	fmt.Println("Workers work Finished")
}