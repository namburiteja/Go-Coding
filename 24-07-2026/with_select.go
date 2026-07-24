package main
import (
	"fmt"
	"time"
)
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func(){
		time.Sleep(8*time.Second)
		ch1<-"idi method 1"
	}()
	go func(){
		time.Sleep(2*time.Second)
		ch2<-"idi method 2"
	}()
	select {
	case msg := <- ch1:
		fmt.Println(msg)
	case msg := <- ch2:
		fmt.Println(msg)
	}
}