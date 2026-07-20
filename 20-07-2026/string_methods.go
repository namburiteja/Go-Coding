package main
import (
	"fmt"
	"strings"
)
func main() {
	s := "        teja   "
	arr := []string{"teja","harika","vandana"}
	fmt.Println(strings.Contains(s,"teja"))
	fmt.Println(strings.HasPrefix("golang", "go"))
	fmt.Println(strings.HasSuffix("photo.png", ".png"))
	fmt.Println(strings.Index("banana", "a"))
	fmt.Println(strings.LastIndex("banana","a"))
	fmt.Println(strings.Count("banana","a"))
	fmt.Println(strings.Repeat("banana ",5))
	fmt.Println(strings.Replace("banana","a","e",2))
	fmt.Println(strings.ReplaceAll("banana","a","e"))
	fmt.Println(strings.Split("a,b,c,d",","))
	fmt.Println(strings.Join(arr,","))
	fmt.Println(strings.TrimSpace("nanna "))
	fmt.Println(strings.Trim("VVVVkingu ra tejaVVV","V"))
	fmt.Println(strings.ToUpper("it is upper"))
	fmt.Println(strings.ToLower("IT IS LOWER"))
	fmt.Println(strings.EqualFold("GO","go")) //check values case insensitive
	fmt.Println(strings.Fields("teja        is          awesome")) //automatically ignores multiple spaces
	fmt.Println(strings.Compare("abc","abb"))
}