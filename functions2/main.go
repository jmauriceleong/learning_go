package main

import (
	"fmt"
)

func salesPerf(){
	items := make(map[string]int)
	items["John"] = 25
	items["Sarah"] = 40
	items["Charlie"] = 30

	for k, v := range items {
		fmt.Printf("%s sold %d items and ", k, v)

		if v < 30 {
			fmt.Println("is below expectations.")
		} else if v >30{
			fmt.Println("is above expectations.")
		}else{
			fmt.Println("meets expectations.")
		}
	}
}

func main(){
	salesPerf()
}