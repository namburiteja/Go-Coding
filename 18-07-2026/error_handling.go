package main
import "fmt"

func divide(a,b int) (int,error) {
	if b==0 {
		return 0, fmt.Errorf("Divide by zero error")
	}else {
		return a/b, nil
	}
}
func main() {
	result, err := divide(10,0)
	if err!=nil {
		fmt.Println(err)
	}else {
		fmt.Println("Result is ", result)
	}
}