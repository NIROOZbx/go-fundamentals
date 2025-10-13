package arithmetic
import "errors"


func Add(a,b int)(int){
	return a+b
}
func Multi(a,b int)(int){
	return a*b
}
func Divide(a,b int)(float64,error){
	if b==0{
		return 0,errors.New("Cant divide by zero")
	}
	return float64(a)/float64(b),nil
}
func Subtract(a,b int)(int){
	return a-b
}