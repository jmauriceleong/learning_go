package main

import "fmt"

func variadic(i ...int) int{
	nums := i
	var s int	

	for x := 0; x < len(nums); x++{
		s += nums[x]
	}
	return s
}

func main(){
	fmt.Println(variadic(54,100,21,3,15))
	i := []int {51,21,31,41,61,71}
	fmt.Printf(" %d\n",variadic(i...))
}