package main
import "fmt"

func main(){
	

	arr1:=[]int{3,4,5,2,1,2}

	for i:=0;i<len(arr1);i++{
		for j:=i+1;j<len(arr1);j++{
			if arr1[i]==arr1[j]{
				arr1=append(arr1[:i],arr1[i+1:]...)
			}
		}
	}
	fmt.Println(arr1)
	

}