package main
import (
	"io"
	"os"
)
func main() {
	src,err := os.Open("read.txt")
	if err!=nil {
		panic(err)
		return
	}
	defer src.Close()
	dest,err := os.Open("tolly.txt")
	if err!=nil {
		panic(err)
		return
	}
	defer dest.Close()
	io.Copy(src,dest)

}