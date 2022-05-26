package main
import (
	"fmt"
)

func main(){
	nums := []int {5,8,9,6,4,7,1,0,2,3}
	// var nums2 []int
	fmt.Println("Before:", nums)

	for i:=0; i < len(nums); i ++ {
		for j:=i+1; j < len(nums); j++{
			if nums[i] > nums[j]{
				nums[i], nums[j] = nums[j], nums[i]
				// nums2 = append(nums2, nums[j])
			}
		}
	}
	fmt.Println("After: ", nums)
}