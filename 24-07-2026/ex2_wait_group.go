package main
import (
	"fmt"
	"sync"
)
func work(i int,wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("worker",i*i)

}
func main() {
	fmt.Println("Now started")
	var wg sync.WaitGroup
	for i:=0;i<100;i++ {
		wg.Add(1)
		go work(i,&wg)
	}
	wg.Wait()
	fmt.Println("Workers work Finished")
}