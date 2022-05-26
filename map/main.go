package main

import (
	// "errors"
	"fmt"
	"os"
	// "reflect"
	// "strconv"
)

// func readArgs() string{
// 	if os.Args[1] == "" {
// 		fmt.Println("You need to enter a key")
// 	}else if reflect.TypeOf(os.Args).String() != "int"{
// 		fmt.Println("You must enter an integer corresponding to a key in the Map")
// 	}
// 	return os.Args[1]
// }

func readMap(names map[string]string)string{
	var matchedstring string
	for k := range names {
		i  := os.Args[1]
		// fmt.Println(i)
		if k == i{
			matchedstring = names[k]
			// fmt.Println(k)
	} else{
		// fmt.Print(err)
	}
	
}
	return matchedstring
}

func main(){
	names := map[string]string{"305":"Sue", "204":"Bob", "631":"Jake", "073":"Tracy"}
	matchedstring := readMap(names)
	fmt.Println("This is the name that matched the key you entered: ", matchedstring)


}