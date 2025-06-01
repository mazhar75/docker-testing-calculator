package mod
import "errors"

func Mod(a int, b int)(int,error){
	if b==0{
		return -1,errors.New("Cann't be divide by zero !")
	}
	return a%b,nil
}