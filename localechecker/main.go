package main

import (
	"fmt"
	"os"
)

type local struct{
	language string
	country string
}

func main(){
	lcollection := []local {
		{"en_US", "United States"},
		{"en_CN", "Canada"},
		{"fr_CN", "Canada"},
		{"fr_FR", "France"},
		{"ru_RU", "Russia"},
	}

	larg := os.Args[1]

	var checklocal = local{larg, ""}
	var check bool

	for i := 0; i < len(lcollection); i ++{
		if checklocal.language == lcollection[i].language {
			check = true
			break
		} else {
			check = false
		}
	}
	if check {fmt.Println("Locale supported: ", larg)} else {fmt.Println("Locale not supported: ", larg)}
}