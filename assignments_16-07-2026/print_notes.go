package main
import (
	"fmt"
)
func main() {
	var n int
	fmt.Print("Enter the amount: ")
	fmt.Scanf("%d",&n)
	no_of_notes := 0
	if n<0 {
		fmt.Println("Not Possible")
		return
	}
	if n>=1000 {
		thousands := n/1000
		n%=1000
		no_of_notes += thousands
		fmt.Println("1000 x",thousands)
	}
	if n>=500 {
		five_hundreds := n/500
		n%=500
		no_of_notes += five_hundreds
		fmt.Println("500 x",five_hundreds)
	}
	if n>=200 {
		two_hundreds := n/200
		n%=200
		no_of_notes += two_hundreds
		fmt.Println("200 x",two_hundreds)
	}
	if n>=100 {
		hundreds := n/100
		n%=100
		no_of_notes += hundreds
		fmt.Println("100 x",hundreds)
	}
	if n>=50 {
		fifty := n/50
		n%=50
		no_of_notes += fifty	
		fmt.Println("50 x",fifty)
	}
	if n>=20 {
		twenty := n/20
		n%=20
		no_of_notes += twenty
		fmt.Println("20 x",twenty)
	}
	if n>=10 {
		ten := n/10
		n%=10
		no_of_notes += ten
		fmt.Println("10 x",ten)
	}
	if n>=5 {
		five := n/5
		n%=5
		no_of_notes += five
		fmt.Println("5 x",five)
	}
	if n>=2 {
		two := n/2
		n%=2
		no_of_notes += two
		fmt.Println("2 x",two)
	}	
	if n>=1 {
		one := n/1
		n%=1
		no_of_notes += one
		fmt.Println("1 x",one)
	}
	if(n==0){
		fmt.Println("No of notes is", no_of_notes)
		return
	}else{
		fmt.Println("Not Possible")
		return
	}

}
