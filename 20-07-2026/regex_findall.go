package main
import (
	"fmt"
	"regexp"
)
func main() {
	re := regexp.MustCompile(`\d+`)
	fmt.Println(re.FindAllString("a 2 b 3 c 4",-1)) 
}