package main
import (
	"fmt"
	"sync"
)
func work(i int,ch chan int,wg *sync.WaitGroup){
	defer wg.Done()
	for job := range ch{
		fmt.Println(i,job)
	}
}
func main() {
	jobs := make(chan int)
	var wg sync.WaitGroup
	for i:=0;i<4;i++ {
		wg.Add(1)
		go work(i,jobs,&wg)
	}
	for i:=0;i<20;i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}