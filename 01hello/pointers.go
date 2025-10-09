package main
import "fmt"

func ChangeNum(num1 *int ,num2 *int) int{
	*num1=100
	*num2=239
	return *num1+*num2

}

func main(){
	num1:=4
	num2:=5

	fmt.Println(ChangeNum(&num1,&num2))
	
}