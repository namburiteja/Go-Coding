package main
import "fmt"
func CheckType(value any) {
	switch v := value.(type) {
		case string:
			fmt.Println("The value is a string:", v)
		case int:
			fmt.Println("The value is an integer:", v)
		case float64:
			fmt.Println("The value is a float:", v)
		case bool:
			fmt.Println("The value is a boolean:", v)
	}

}
func main() {
	CheckType("Hello, World!")
	CheckType(42)
	CheckType(3.14)
	CheckType(true)
}