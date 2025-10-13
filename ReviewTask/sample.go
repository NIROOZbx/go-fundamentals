package main 
import "fmt"
import "errors"
// import "bufio"
// import "os"


// // arr:=[4]int{1,2,3,5}
// // slice:=[]int{1,2,3,4,5}

// // m1:=map[string]int{Name:"nirooz",age:20}


// // for i:=0;i<5;i++{
// // 	fmt.Println(i)
// // }

// // for i,_:=range slice{
// // 	fmt.Println(slice[i])
// // }

// // type Person struct{
// // 	Name int
// // 	Age int
// // 	Address string
// // }

// // p1:=Person{Name:"sfds",Age:43,Address:"Kozhikode"}

// // num1:=10

// // if num1<10{
// // 	fmt.Println("Num1 is less than 10")
// // }else{
// // 	fmt.Println("Num is greate than 10")
// // }

// func Calc(num1,num2 int,input string){

// 	fmt.Println(input)


// 	switch input{

// 	case "Addition":
		
// 		fmt.Println(num1+num2) 
		
// 	case "Subtraction":
	
// 		fmt.Println(num1-num2) 
		
// 	case "Divison":
// 		fmt.Println(num1/num2) 
		
// 	case "Multiplication": 
// 		fmt.Println(num1*num2) 

// 	default:
// 		fmt.Println("Enter a valid operation")
// 	}


// }


// func main(){
	

// 	reader:=bufio.NewReader(os.Stdin)
// 	fmt.Println("Choose one")
// 	fmt.Println("Addition")
// 	fmt.Println("Subtraction")
// 	fmt.Println("Multiplication")
// 	fmt.Println("Divison")
// 	fmt.Println("----------------")

	

// 	input,_:=reader.ReadString('\n')

// 	num1,num2:=5,6

// 	fmt.Println(Calc(num1,num2,input))


	


// }

func DivisonByZero(a,b int) (int,error){

	if b==0{
		return 0,errors.New("Division by zero")
	}else{
		return a/b,nil
	}

}


func main(){

	result,err:=DivisonByZero(5,0)

	if err==nil{
		fmt.Println(result)
	}else{
		fmt.Println(err)
	}

}