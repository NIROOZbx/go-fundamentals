package main

import "fmt"

var students=map[string]int{"har":56}


func main(){
	students["john"]=50
students["alex"]=87
students["david"]=99
students["Alice"]=89
students["Dave"] = 92
students["Alice"] = 95
	
	var val,exists=students["alex"]

	if exists{
		fmt.Println("Key exists with value",val)
	}else{
		fmt.Println("Cant find the key")
	}

	delete(students,"Alice")

	fmt.Println(students)
}