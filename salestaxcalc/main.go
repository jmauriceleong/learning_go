package main
import (
	"fmt"

)

type goods struct{
	item string
	cost float32
	salestax float32
}

func salestaxcalc(cost float32, salestax float32) float32 {
	
	var salestaxtotal float32
	salestaxtotal = cost * salestax
	return salestaxtotal

}

func main(){
	allgoods := []goods {
	
	goods {
		item: "cake",
		cost: 0.99,
		salestax: 0.075,
	},
	goods{
		item: "milk",
		cost: 2.75,
		salestax: 0.015,
	},
	goods{
		item: "butter",
		cost: 2.75,
		salestax: 0.015,
	},
	
	}
	var totaltax float32
	for i := 0; i < len(allgoods); i++{
		totaltax  += salestaxcalc(allgoods[i].cost,  allgoods[i].salestax)
	}

	fmt.Println("Sales Tax Total :", totaltax)
}