package main
import (
	"fmt"
	"regexp"
)
func main() {
	re := regexp.MustCompile(`,`)
fmt.Println(re.Split("java,python,go",-1))
}