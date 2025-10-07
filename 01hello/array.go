package main
import "fmt"

func findMinMax()(int,int){
	var nums=[7]int{3,2,9,24,1,4,5}

	min:=nums[0]
	max:=nums[0]

	for i:=0;i<len(nums);i++{
		if min>nums[i]{
			min=nums[i]
		}
		if max<nums[i]{
			max=nums[i]
		}
	}

	return max,min


}

func main(){
	
	 fmt.Println(findMinMax())
}