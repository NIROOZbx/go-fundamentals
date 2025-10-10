package main

import "fmt"
import "time"

func Odd(num int){
	if num%2!=0{
		fmt.Println("Odd")
	}else{
		fmt.Println("Even")
	}
	

}

func Even(num int){
	if num%2==0{
		fmt.Println("Even")
	}else{
		fmt.Println("Odd")
	}

}

func main(){
	go Odd(5)
	go Even(8)
	time.Sleep(1000*time.Millisecond)
	fmt.Println("Waited 1 second")

}