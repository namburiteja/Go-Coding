package main
import (
	"fmt"
	"context"
	"time"
)
func worker(ctx context.Context) {
	for{
		select {
		case <-ctx.Done():
			fmt.Println("Worker stopped")
			return
		default :
			fmt.Println("Working")
			time.Sleep(time.Second)
		}
	}
}
func main() {
	ctx,cancel := context.WithCancel(context.Background())
	go worker(ctx)
	time.Sleep(3*time.Second)
	cancel()
	time.Sleep(time.Second)
}