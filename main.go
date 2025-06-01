package main
import (
	"fmt"
	"example.com/add"
	"example.com/sub"
	"example.com/mul"
	"example.com/div"
	"example.com/mod"
)
func main(){
	 var a,b int
	 var s string
	 fmt.Println("Input your operands")
	 fmt.Scan(&a,&b)
	 fmt.Println("enter your operator")
     fmt.Scan(&s)
	 op := rune(s[0])
	 switch op{
	   case '+': fmt.Println(add.Add(a,b))
	   case '-': fmt.Println(sub.Sub(a,b))
	   case '*': fmt.Println(mul.Mul(a,b))
	   case '/': res,err := div.Div(a,b)
	             if err==nil{
					fmt.Println(res)
				 } else {
					fmt.Println(err)
				 }
	   case '%': res,err := mod.Mod(a,b)
	             if err==nil{
					fmt.Println(res)
				 } else {
					fmt.Println(err)
				 }
	    default:fmt.Println("Invalid operator")			 
	 }


}