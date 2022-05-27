package main
import (
	"fmt"
)

func fizzBuzz(start, end int){
	fmt.Println("FizzBuzz is in Control")
	for i := start; i < end; i ++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		}else if i%3 == 0{
			fmt.Println("Fizz")
		} else if i%5 ==0 {
			fmt.Println("Buzz")
		}else {
			fmt.Println(i)
		}
	}
}

func main(){
	fmt.Println("Main is in control")
	fizzBuzz(25,100)
	fmt.Println("Main is back in control")
}