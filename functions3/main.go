package main

import (
	"fmt"
	// "os"
	// "strconv"
)

func fizzBuzz(i int) (int, string){
	switch{
	case i%15 == 0:
		return i, "FizzBuzz"
	case i%3 == 0:
		return i, "Fizz"
	case i%5 == 0:
		return i, "Buzz"
	default:
		return i, "neither Fizz nor Buzz"
	}
}

func main(){

	n, s := fizzBuzz(35)
	fmt.Printf("%d is %s\n", n, s)
}