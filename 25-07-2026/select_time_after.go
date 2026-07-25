package main

import (
	"fmt"
	"time"
)

func main(){

	ch:=make(chan string)

	select{

	case msg:=<-ch:

		fmt.Println(msg)

	case <-time.After(3*time.Second):

		fmt.Println("Timeout")

	}
}