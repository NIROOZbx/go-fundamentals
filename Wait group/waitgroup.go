package main

import "fmt"
import "sync"

func Increment(wg *sync.WaitGroup,mu *sync.Mutex,count *int) {
	defer wg.Done()
	mu.Lock()
	 *count++
	 defer mu.Unlock()
	 fmt.Println(*count) 

}

func main(){

	var wg sync.WaitGroup
	var mu sync.Mutex
	var counter int

	for i:=0;i<100;i++{
		wg.Add(1)
		go Increment(&wg,&mu,&counter)

	}
fmt.Println("Starting the workers")
wg.Wait()
fmt.Println("Done with the task")
}