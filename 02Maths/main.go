package main

import "fmt"

func Calc( a , b int)  (int, int, float64, int){

	sum:=a+b
	multi:=a*b
	division:=float64(a)/float64(b)
	subtract:=a-b

	return sum,multi,division,subtract
}

func main() {

	q,w,e,r:=Calc(5,6)
	q=q+w
	fmt.Println(q,w,e,r)
     
}
