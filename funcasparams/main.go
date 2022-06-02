package main

import (
	"fmt"
	
)

type calc func(int, int) int

func main(){
	calculator(add, 2, 3)
	calculator(subtract, 10, 7)
}
func add(i, j int)int{
	return i + j	
}
func subtract(i, j int)int{
	return i - j
}
func calculator(f calc, i, j int){
	fmt.Printf("The result of the operation is %d\n", f(i,j))
}