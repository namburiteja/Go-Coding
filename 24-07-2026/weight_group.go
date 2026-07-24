package main
import (
	"fmt"
	"sync"
)
func hello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello Wait Group")
}
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go hello(&wg)
	wg.Wait()
	fmt.Println("OK")
}