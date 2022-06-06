package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("Main Function")
	
	go countNumbers(20)

	time.Sleep(1 * time.Second)

	fmt.Println("End main function")
}

func countNumbers(limit int){
	num := 0

	for i := 1; i < limit; i++{
		num += i
	}

	fmt.Println("Num:", num)
}