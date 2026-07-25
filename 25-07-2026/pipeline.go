package main

import "fmt"

func multiply(in <-chan int) <-chan int {

	out := make(chan int)

	go func() {

		defer close(out)

		for n := range in {

			out <- n * 2
		}

	}()

	return out
}

func square(in <-chan int) <-chan int {

	out := make(chan int)

	go func() {

		defer close(out)

		for n := range in {

			out <- n * n
		}

	}()

	return out
}

func main() {

	input := make(chan int)

	go func() {

		defer close(input)

		for i := 1; i <= 5; i++ {

			input <- i
		}
	}()

	stage1 := multiply(input)

	stage2 := square(stage1)

	for value := range stage2 {

		fmt.Println(value)
	}
}