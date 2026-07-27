package main
import (
	"fmt"
	"sync"
) 
var counter int
var mu sync.Mutex
func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	counter++
	mu.Unlock()
}
func main() {
	var wg sync.WaitGroup
	for i:=1;i<=100;i++ {
		wg.Add(1)
		go increment(&wg)
	}
	wg.Wait()
	fmt.Println(counter)
}