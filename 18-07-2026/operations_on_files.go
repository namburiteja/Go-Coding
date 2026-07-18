package main
import (
	"fmt"
	"os"
)
func readallfiles() {
	files , _ := os.ReadDir(".")
	for _,file := range files {
		fmt.Println(file)
	}
}
func readOnlyDirectories() {
	files,_ := os.ReadDir(".")
	for _,file := range files {
		if file.IsDir() {
			fmt.Println(file.Name())
		}
	}
}
func main() {
	// readallfiles()
	readOnlyDirectories()
}