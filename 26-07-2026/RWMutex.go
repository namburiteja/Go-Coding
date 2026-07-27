package main
import (
	"fmt"
	"sync"
)
var counter int
var mu sync.RWMutex
func read(wg *sync.WaitGroup) {
	defer wg.Done()
	mu.RLock()
	fmt.Println(counter)
	mu.RUnlock()
}
func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	counter++
	mu.Unlock()
}
func main(){
	var wg sync.WaitGroup
	for i:=0;i<20;i++ {
		wg.Add(1)
		go increment(&wg)
		go read(&wg)
	}
	wg.Wait()
}