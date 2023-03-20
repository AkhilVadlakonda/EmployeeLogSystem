package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// int
	var TotalWorkingdays int = 5
	var Activelogged int = 3
	var payperday int64 = 150
	y := int(payperday)
	var bonus int = rand.Intn(50) + 1
	fmt.Printf("Employee Abscence is %d days\n", TotalWorkingdays-Activelogged)
	fmt.Printf("Pay for the week = %d\n\n\n", y*(TotalWorkingdays-Activelogged)+bonus)

	//time data type
	Clockin := time.Date(2023, time.March, 10, 9, 30, 0, 0, time.UTC)
	Clockout := time.Date(2023, time.March, 10, 12, 0, 0, 0, time.UTC)
	duration := Clockout.Sub(Clockin)
	fmt.Printf("Active hours: %v\n", duration)
	now := time.Now()
	timeStr := now.Format("2006-01-02 15:04:05")
	fmt.Println("Current time:", timeStr, "\n")

	// float64 data type
	var totalhours float64 = 40.4777
	var pay float64 = 20.334
	var deductions float64 = 20.66
	round := math.Round(totalhours*100) / 100
	absded := math.Abs(deductions)
	fmt.Println("Rounded hours:", round)
	fmt.Printf("Salary = %f\n", round*pay)
	fmt.Printf("Final pay = %f\n\n", (totalhours*pay)-absded)

	// string data type
	var EmployeeName string = "Rachana Goli"
	fmt.Printf("The length of the string \"%s\" : %d\n", EmployeeName, len(EmployeeName))
	fmt.Printf("Editing Last Name of \"%s\" to \"Reddy\" Edited Name: \"%s\"\n", EmployeeName, strings.Replace(EmployeeName, "Goli", "Reddy", 1))
}
