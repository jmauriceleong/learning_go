package main

import (
	"fmt"
	"unicode"
)
// Online portal creates user accounts and accept passwords with the following criteria:
	// Passwords are to be 8-15 characters long
	// Have a lowercase letter
	// Have an uppercase letter
	// Have a number
	// Have a symbol
func passwordChecker(pw string) bool {
	pwR := []rune(pw)

	if len(pwR) <8 {
		return false
	} 
	if len(pwR) > 15 {
		return false
	}

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSymbol := false

	for _, v := range pwR{
		if unicode.IsUpper(v){
			hasUpper = true
		} 

		if unicode.IsLower(v){
			hasLower = true
		}
		if unicode.IsNumber(v){
			hasNumber = true
		}
		if unicode.IsSymbol(v) || unicode.IsPunct(v){
			hasSymbol = true
		}
	}

	switch {
	case !hasUpper:
		fmt.Println("Your password must have at least 1 uppercase letter!")
	case !hasLower:
		fmt.Println("Your password must have at least 1 lowercase letter!")
	case !hasNumber:
		fmt.Println("Your password must have at least 1 number!")
	case !hasSymbol:
		fmt.Println("Your password must have at least 1 symbol!")
	case len(pwR) <8:
		fmt.Println("Your password must be at least 8 characters long!")
	case len(pwR) >15:
		fmt.Println("Your password cannot be longer than 15 characters!")
	}

	return hasLower && hasNumber && hasSymbol && hasUpper
}

func main(){
	if passwordChecker("akHsa0dalhfaslf") {
		fmt.Println("password good")
		} else {
		fmt.Println("password bad")
	}
}