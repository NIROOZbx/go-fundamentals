package main
import "fmt"
import "sync"
import "time"

func main(){

	ch:=make(chan string)
	var wg sync.WaitGroup


	wg.Add(3)
	go func(){ 
		defer wg.Done()
		ch <-"First"
		ch<-"Fourth"
		}()
	go func(){ 
		defer wg.Done()
		ch <-"Second"
		}()
	go func(){ 
		defer wg.Done()
		ch <-"Third"
		}()
		go func(){
			wg.Wait()
			close(ch)
		}()
		
		for Message:=range ch{
			fmt.Println(Message)
			time.Sleep(1000 * time.Millisecond)
		}

	fmt.Print("Ended listening and recieving")

}