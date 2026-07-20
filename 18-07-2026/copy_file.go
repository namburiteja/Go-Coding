package main

import (
	"io"
	"os"
)

func main() {

	src, err := os.Open("read.txt")
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create("copy.txt")
	if err != nil {
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return
	}
}