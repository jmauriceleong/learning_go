package main
import (
	"fmt"
)
type Weekday int

const (
	Sunday Weekday = iota 
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)
func (w Weekday) day(i int)string{
	var s string
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	for x:=0; x < len(days); x++{
		if x == i{
			s = days[x]
		}
	}
	return s
}


type Employee struct {
	Id int
	FirstName string
	LastName string
}

type Developer struct {
	emp Employee
	HourlyRate int
	Workweek [7]int
}

func (d *Developer) LogHours(w Weekday, h int){
	for i := 0; i < len(d.Workweek); i ++ {
		if int(w) == i {
			d.Workweek[i] = h
		}
	}
}

func (d *Developer) HoursWorked() int{
	var s int
	for i := 0; i < len(d.Workweek); i++{
		s += d.Workweek[i]
	}
	return s
}

func main(){
	var d Developer
	d.emp.Id =1
	d.emp.FirstName = "Arthur"
	d.emp.LastName = "Ashe"
	d.HourlyRate = 40
	
	var m = Monday
	var w = Tuesday
	d.LogHours(m, 10)
	d.LogHours(w, 12)
	fmt.Printf("Hours worked on %s: %d\n", m.day(int(m)), d.Workweek[m])
	fmt.Printf("Hours worked on %s: %d\n", w.day(int(w)), d.Workweek[w])
	fmt.Printf("Total hours for employee %s %s is : %d\n", d.emp.FirstName, d.emp.LastName, d.HoursWorked())
}