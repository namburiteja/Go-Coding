package main
import (
	"fmt"
)
func download(id int,jobs chan int){
	for i:=id;i<=id+5;i++ {
		jobs <- i
	}
}
func main() {
	jobs := make(chan int)
	go download(1,jobs)
	go download(100,jobs)
	for job := range jobs {
		fmt.Println(job)
	}
}