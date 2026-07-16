package main
import "fmt"

func main(){
	var n,k int
	fmt.Print("Enter two numbers: ")
	fmt.Scanf("%d,%d",&n,&k)
	arr := []int{}
	for i:=1;i<=n;i++ {
		if n%i == 0 {
			arr = append(arr,i)
		}
	}
	fmt.Println(arr)
	int length := len(arr)
	if(length < k) {
		fmt.Println(1)
	}else {
		fmt.Println(arr[(length-k)])
	}

}