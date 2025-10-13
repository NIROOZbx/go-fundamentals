package main

import "fmt"

var studentss=map[string]int{"har":56}


func main(){
	studentss["john"]=50
studentss["alex"]=87
studentss["david"]=99
studentss["Alice"]=89
studentss["Dave"] = 92
studentss["Alice"] = 95
	
	var val,exists=studentss["alex"]

	if exists{
		fmt.Println("Key exists with value",val)
	}else{
		fmt.Println("Cant find the key")
	}

	delete(studentss,"Alice")

	fmt.Println(studentss)
}