package main
import "fmt"

func main(){
	ch:=make(chan string ,2)

	fmt.Println("Send the data")
	ch<-"Hello"
	ch<-"World"
	ch<-"sf"

	fmt.Println("Recieving the data")

	ch1:=<-ch
	ch2:=<-ch
	ch3:=<-ch

	fmt.Println("Got the data",ch1,ch2,ch3)


	
}
