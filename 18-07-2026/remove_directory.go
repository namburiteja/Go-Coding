package main 
import ( "fmt" 
"os")
func main () {
	err := os.Remove("assignments_renamed")
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println("Directory removed successfull")
	}
}