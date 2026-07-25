package main
import (
	"fmt"
	"time"
)
func main() {
	ticker := time.NewTicker(2*time.Second)
	timer := time.NewTimer(6*time.Second)
	for {
		select {
		case <- ticker.C :
			fmt.Println("Beat it")
		case <- timer.C :
			ticker.Stop()
			fmt.Println("stopping")
			return
		}
	}
}