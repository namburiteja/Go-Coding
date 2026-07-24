package main

import "fmt"

func main() {

	ch := make(chan int)

	go func() {

		fmt.Println("Sending")

		ch <- 50

		fmt.Println("Sent")

	}()

	fmt.Println("Receiving")

	value := <-ch

	fmt.Println(value)
}