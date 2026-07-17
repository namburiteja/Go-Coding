package main
import "fmt"

func main() {
	var value any = "Hello, World!"

	number := value.(string) 

	fmt.Println(number + "5")
}