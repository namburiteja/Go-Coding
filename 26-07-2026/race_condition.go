package main

import (
	"fmt"
	"sync"
)

var counter int

func increment(wg *sync.WaitGroup) {
	defer wg.Done()

	counter++
}

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {

		wg.Add(1)

		go increment(&wg)
	}

	wg.Wait()

	fmt.Println(counter)
}