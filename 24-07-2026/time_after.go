package main

import (
	"fmt"
	"time"
)

func main(){

	fmt.Println("Waiting...")

	<-time.After(3*time.Second)

	fmt.Println("Finished")
}