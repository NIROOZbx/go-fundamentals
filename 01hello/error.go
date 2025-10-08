package main

import "fmt"
import "errors"

func divideByZero(a,b int) (int,error){

	if b==0{
		return 0,errors.New("Cant divide by zero")
	}

	return a/b,nil

}

func main(){
  result,err:=divideByZero(5,0)

  if(err==nil){
	fmt.Println(result)
  }else{
	fmt.Println(err)
  }
}