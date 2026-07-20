package main
import (
	"fmt"
	"regexp"
)
func main() {
	re := regexp.MustCompile(`^\d{10}$`)
	fmt.Println(re.MatchString("9876543210"))
}