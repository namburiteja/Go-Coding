package main
import (
	"fmt"
	"regexp"
)
func main() {
	re := regexp.MustCompile(`\d+`)
	fmt.Println(re.FindString("Age 2 kay 23"))
}