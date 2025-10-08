package main

import "fmt"

var Exmp=func (a,b int, c ...string)(int,int,string, string){
	return a,b,c[0],c[1]
}


func main(){
	q,w,e,r:=Exmp(2,4,"nirooz","welcome")
	fmt.Println(q,w,e,r)



}