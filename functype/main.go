package main

import (
	"fmt"
)

func main(){
	devSalary := salary(50,2000, developerSalary)
	bossSalary := salary(150000, 25000, managerSalary)
	fmt.Printf("Boss Salary: %d\n", bossSalary)
	fmt.Printf("Developer's Salary: %d\n", devSalary)

}

func salary(x,y int, f func(int,int)int)int{
	pay := f(x,y)
	return pay
}

func managerSalary(baseSalary, bonus int)int{
	return baseSalary + bonus
}

func developerSalary(hourlyrate, hoursWorked int)int{
	return hourlyrate * hoursWorked
}

func nonLoggedHours() func(int)int{
	
}