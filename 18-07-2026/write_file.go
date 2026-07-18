package main
import (
	"os"
)
func main() {
	data := []byte("Master The Blaster")
	err := os.WriteFile("read.txt",data,0644)
	if err!=nil {
		panic(err)
	}
	return
}