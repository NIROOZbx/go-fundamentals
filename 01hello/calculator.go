package main 
import "fmt"


func main(){
    var num1,num2 int
	var operation string
	fmt.Println("Enter num1")
	fmt.Scanln(&num1)
	fmt.Println("Enter num2")
	fmt.Scanln(&num2)
    fmt.Println("----------------")
	fmt.Println("Addition")
	fmt.Println("Subtraction")
	fmt.Println("Multiplication")
	fmt.Println("Division")
	fmt.Println("----------------")
	fmt.Scanln(&operation)

	switch operation{
		case "Addition":
		
		fmt.Println(num1+num2) 
		
	case "Subtraction":
	
		fmt.Println(num1-num2) 
		
	case "Division":
		if num2==0{
			fmt.Println("cannot divide by zero")
		}else{
		fmt.Println(float64(num1)/float64(num2)) 
		 }
		
	case "Multiplication": 
		fmt.Println(num1*num2) 

	default:
		fmt.Println("Enter a valid operation")
	}




}