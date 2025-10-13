package main
import "fmt"

type Person struct{
	Name string
	Address string
	Age int
	Cars []string
}

func (p Person) Alldetails() {
	fmt.Printf("Fulll details%+v",p)
}


func main(){

	p1:=Person{Name:"Nirooz",Address:"kozhikode",Age:21,Cars:[]string{"honda","buugatti"}}
	p1.Name="John"
	p1.Alldetails()


}