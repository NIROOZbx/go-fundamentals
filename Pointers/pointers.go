package main
import "fmt"

type CustomError struct{
	Message string
	Status int
}

func (c *CustomError) Error() string {
	return fmt.Sprintf("No Error %s %d",c.Message,c.Status)
}

func GetError() error{
	return &CustomError{Message:"INvalid input",Status:404}
}

func ChangeNum(num1 *int ,num2 *int) int{
	*num1=100
	*num2=239
	&num1=844
	return *num1+*num2

}

func main(){
	num1:=4
	num2:=5

	e1:=GetError()

	e1,ok:=e1.(*CustomError)
	fmt.Println(e1,ok)



	

	fmt.Println(ChangeNum(&num1,&num2))
	
}