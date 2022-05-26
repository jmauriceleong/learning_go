package main

import (
	"fmt"
	"reflect"
)
type alltypes struct{
	num int
	fl  float32
	str string
	bl  bool
	ukn interface{}
}

func acceptTypes(v interface{}) string {

	switch t := v.(type){
	case int, int32, int64:
		return fmt.Sprint("The type of this parameter {} is int.", t)
	case float32, float64:
		return fmt.Sprint("The type of this parameter {} is float.", t)
	case bool:
		return fmt.Sprint("The type of this parameter {} is bool.", t)
	case string:
		return fmt.Sprint("The type of this parameter {} is string.", t)
	default:
		return fmt.Sprint("The type of this parameter {} is UNKNOWN.", t)
	}
}

func main(){
	var ukn interface{}
	al := alltypes{5, 1.23556, "Charles", true, ukn}
	
	v := reflect.ValueOf(al)

	for i := 0; i < v.NumField(); i++{
		fmt.Println(v.Field(i))
		fmt.Println(acceptTypes(reflect.Value.Type(v.Field(i))))
	}
}

