package main
import "fmt"

func change(n *int) {
	*n = 20
}

func main() {
	n := 10
	change(&n)
	fmt.Println(n)
	//another way to print n
	x := new(int)
	*x = 18
	fmt.Println(*x)
}