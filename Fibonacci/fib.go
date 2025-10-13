package main

import "fmt"
func inter(val interface{}) interface{}{
	return val
}

func main() {
//  a,b:=0,1
// 	for i:=0;i<9;i++ {
// 		temp:=a
// 		a=b
// 		b=temp+a
// 		fmt.Println(a)
// 	}

	fmt.Println(inter("hello"))

}