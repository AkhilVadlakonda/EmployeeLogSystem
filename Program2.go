package main

import "fmt"

func main() {

	// Struct
	type Employee struct {
		EmpName     string
		Age         int
		Designation string
	}
	var emp1 Employee = Employee{"Rachana Goli", 25, "Manager"}
	fmt.Println("Employee:", emp1)

	// Array
	var hours [5]int = [5]int{3, 5, 3, 6, 8}
	fmt.Println("working hours in a week:", hours)

	// For loop
	var Totalhours int = 0
	for i := 0; i < len(hours); i++ {
		Totalhours = Totalhours + hours[i]
	}
	fmt.Printf("Total hours for week: %d ", Totalhours)
	fmt.Println()

	// If-else
	if Totalhours > 20 {
		fmt.Println(emp1.EmpName, "is a full-time", emp1.Designation)
	} else {
		fmt.Println(emp1.EmpName, "is a part-time", emp1.Designation)
	}
}
