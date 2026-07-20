package main
import (
	"fmt"
	"regexp"
)
func main() {
	re := regexp.MustCompile(`cat`)
	fmt.Println(re.ReplaceAllString("catjink vha cat cat ringa","lion"))
}