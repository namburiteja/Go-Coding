package main
import (
	"fmt"
	"strconv"
)
func main(){
	fmt.Println(strconv.Atoi("12a3"))
	fmt.Println(strconv.Itoa(1))
	fmt.Println(strconv.ParseInt("420",10,32)) //string,base,32 or 64 bit
	fmt.Println(strconv.ParseFloat("12.25",64)) //string,32 or 64 bit
	fmt.Println(strconv.ParseBool("1")) // can write - 0,1,true,false,t,f
	fmt.Println(strconv.FormatInt(420,10)) // int,base
	fmt.Println(strconv.FormatFloat(12.25,'f',2,64)) // float,byte,upto point,float64
	fmt.Println(strconv.FormatBool(true)) //true or flase can write

}