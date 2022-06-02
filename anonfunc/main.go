package main

import (
	"fmt"
)


func main(){
	message := "Greeting"
	
	func(str string)  {
		fmt.Println(str)
	}(message)

	f := func(str string){
		fmt.Println(str)
	}

	f(message)
}