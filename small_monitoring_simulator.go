package main
import (
	"fmt"
	"time"
)
func main() {
	ticker := time.MakeTicker(500*time.Millisecond)
	defer ticker.Stop()
	timer := time.NewTimer(5*time.Second)
	
}