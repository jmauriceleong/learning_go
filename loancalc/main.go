package main

import (
	"fmt"
	"errors"
	"math"
)

type applicant struct {
	income float32
	cscore int
	amt float32
	term int
}
//define struct to hold all loan details
type loandetails struct{
	creditscore int
	monthlyincome float32
	loanamount float32
	loanterm int
	monthlypayment float32
	interestrate float32
	totalcost float32
	approved bool	
}

func processloan(income float32, creditscore int, loanamt float32, loanterm int) loandetails{
	var rate float32
	var total float32
	var errArray []error
	var approval bool
	
	// determine interest rate based on credit score
	if creditscore < 450 {
		rate = 0.2
	}else {
		rate = 0.15
	}
	
	// more error checking in the next 2 "if" blocks
	if loanterm % 12 > 0 {
		errArray = append(errArray, errors.New("this loan term is not divisible by 12"))
	}
	if income < 0 || creditscore < 0 || loanamt < 0 || loanterm < 0{
		errArray = append(errArray, errors.New("invalid credit score income loan amount or loan term"))
	} 
	
	// Calculate repayment
	repayment := loanamt*(rate/12*(float32(math.Pow(float64(1+(rate/12)),float64(loanterm)))))/float32((math.Pow(float64(1+(rate/12)),float64(loanterm))-1))
	// Calculate total amount payable over the term of the loan
	total = repayment * float32(loanterm)
	// Apply more criteria, if false store errors in an array
	if creditscore < 450{
		if repayment > (income * 0.1){
			errArray = append(errArray, errors.New("your repayment is too high"))
		}
	}else {
		if repayment > (income * 0.2){
			errArray = append(errArray, errors.New("your repayment is too high"))
		}
	}
	// If error array is not empty, set approved status to false, otherwise true if empty
	if len(errArray) > 0 {
		approval = false
	}else{
		approval = true
	}
	// Loop through error array and display errors
	for j := 0; j < len(errArray); j++{
		fmt.Println(errArray[j])
	}
	// Return a loan details struct with all the relevant info to display
	return loandetails{
		creditscore: creditscore,
		monthlyincome: income,
		loanamount: loanamt,
		loanterm: loanterm,
		monthlypayment: repayment,
		interestrate: rate*100,
		totalcost: total,
		approved: approval,
	}
}

func main(){
	// Declare 2 applicant structs
	applicant1 := applicant{5500.00, 600, 10000, 6}
	applicant2 := applicant{6500.00, 300, 9000, 36}

	// call the process loan function on each applicant and store in 2 seperate variables
	app1outcome := processloan(applicant1.income, applicant1.cscore, applicant1.amt, applicant1.term)
	app2outcome := processloan(applicant2.income, applicant2.cscore, applicant2.amt, applicant2.term)

	// Print out the outcome for each applicant using the struct returned from the process loan function
	// applicant 1
	fmt.Println("Applicant 1\n -----------\nCredit Score:\t", app1outcome.creditscore,
	 "\nIncome:\t\t", app1outcome.monthlyincome,
	 "\nLoan Amount:\t", app1outcome.loanamount,
	 "\nLoan Term:\t",app1outcome.loanterm,
	 "\nMonthly Payment:", app1outcome.monthlypayment,
	 "\nRate:\t\t", app1outcome.interestrate,
	 "\nTotal Cost:\t", app1outcome.totalcost,
	 "\nApproved:\t",app1outcome.approved,
	)
	// applicant 2
	fmt.Println("\nApplicant 2\n -----------\nCredit Score:\t", app2outcome.creditscore,
	 "\nIncome:\t\t", app2outcome.monthlyincome,
	 "\nLoan Amount:\t", app2outcome.loanamount,
	 "\nLoan Term:\t",app2outcome.loanterm,
	 "\nMonthly Payment:", app2outcome.monthlypayment,
	 "\nRate:\t\t", app2outcome.interestrate,
	 "\nTotal Cost:\t", app2outcome.totalcost,
	 "\nApproved:\t",app2outcome.approved,
	)
	
}