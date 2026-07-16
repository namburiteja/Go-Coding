package main
import (
	"fmt"
)
func main() {
	n :=1100
	no_of_notes := 0
	thousands :=0
	hundreds :=0
	two_hundreds :=0
	five_hundreds :=0
	if n<100 {
		fmt.Println("No of notes is 0")
	}
	if n>=1000 {
		thousands = n/1000
		n%=1000
	}
	if n>=500 {
		five_hundreds = n/500
		n%=500
	}
	if n>=200 {
		two_hundreds = n/200
		n%=200
	}
	if n>=100 {
		hundreds = n/100
		n%=100
	}
	if(n==0){
		no_of_notes = (thousands*10) + (five_hundreds*5) + (two_hundreds*2) + hundreds
		fmt.Println("No of notes is", no_of_notes)
		return
	}else{
		fmt.Println("Not Possible")
		return
	}

}