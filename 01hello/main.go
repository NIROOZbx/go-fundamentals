package main

import "fmt"

func main() {
 a,b:=0,1
	for i:=0;i<9;i++ {
		temp:=a
		a=b
		b=temp+a
		fmt.Println(a)
	}

}
