package main
import (
	"os"
)
func main() {
	_,err := os.ReadFile("../18-07-2026/reading.txt")
	if err!=nil {
		panic("There is no such file")
		return
	}
}