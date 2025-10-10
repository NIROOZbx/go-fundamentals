package main
import "fmt"

func main(){

	ch:=make(chan string)

	fmt.Println("Starting the main function")

	go func(){
		fmt.Println("Starting the go routine")

		ch <-"Go routine work is done"

	}()

		ch1:=<-ch

		fmt.Println("Recived the message\n",ch1)
}