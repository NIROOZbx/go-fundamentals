package main
import "fmt"

func main(){
	 cars:= []string{"bugatti","lamborghini","honda","toyota"}
	 cars=append(cars,"Mustang")

	for i:=0; i<len(cars); i++ {
		fmt.Println(cars[i])
		
	}
}