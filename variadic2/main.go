package main

import(
	"fmt"
	"math"
)

func main(){
	j := 9

	x := func(i int)int{
		return i * i
	}

	y := func(i int)float64{
		return math.Sqrt(float64(x(j)))
	}

	fmt.Printf("The square of %d is %d\n", j, x(j))
	fmt.Printf("The square root of %d is %f\n", x(j), y(x(j)))
}