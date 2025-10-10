package main

import "fmt"
import "sync"


func Worker(id int , wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Printf("Worker %d: starting task\n", id)

}

func main(){
var wg sync.WaitGroup

for i:=0;i<5;i++{
	wg.Add(1)
	go Worker(i, &wg)
}

fmt.Println("Main: Waiting for all workers to finish...")
	wg.Wait()
	fmt.Println("Main: All workers have completed their tasks.")
}