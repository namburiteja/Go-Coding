package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d processing Job %d\n", id, job)
	}
}

func main() {

	jobs := make(chan int)

	var wg sync.WaitGroup

	// Create 3 workers
	for i := 1; i <= 3; i++ {

		wg.Add(1)

		go worker(i, jobs, &wg)
	}
	// Send 10 jobs
	for job := 1; job <= 10; job++ {

		jobs <- job
	}

	close(jobs)

	wg.Wait()	
}