package main
import (
	"strings"
	"fmt"
)
func main() {
	var b strings.Builder
	b.WriteString("Hello")
	b.WriteString(" ")
	b.WriteString("World")
	fmt.Println(b.String())
}