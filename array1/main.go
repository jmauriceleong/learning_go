package main

import (
	"fmt"
)

func my_arr(args ...int) int {
	sum := 0
	for _, v := range args {
		sum = sum + v
	}
	return sum
}

func main() {
	arr := []int{2, 4, 5}
	sum := my_arr(arr...)
	fmt.Println("Sum is ", sum)

	arr2 := [...]string{"I", "am", "a", "little", "boy"}
	for i := 0; i < len(arr2); i++ {
	fullstring := fmt.Sprint(arr2[i]," ")
	fmt.Print(fullstring)
	
	}
	fmt.Println()
	var arr3 [10]int

	for j := 0; j < len(arr3); j++{
		arr3[j] = j + 1 
	}

	fmt.Println(arr3)

}
