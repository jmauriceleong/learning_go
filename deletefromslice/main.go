package main

import (
	"fmt"
	"os"
	"strconv"
)

// implement function to delete an element from a slice
func deletefromslice(nums []int,el int) []int{
	 
	 var nums2 []int
	 for i := 0; i < len(nums); i ++{
		 if nums[i] == int(el){
			 nums[i] = 0
		 } else{
			 fmt.Println("Sorry, cannot find this number!")
			 break
		 }

		 if nums[i] != 0 {
			 nums2 = append(nums2, nums[i])
		 }
	 }
	 return nums2
}

func main(){
	el, _ := strconv.Atoi(os.Args[1])
	fmt.Println(el)
	nums := []int {100, 204, 300, 501}
	fmt.Println("Array before delete: ", nums)

	fmt.Println("Array after delete: ", deletefromslice(nums,el))
}