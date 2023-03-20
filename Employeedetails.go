package main

import (
	"fmt"
	"sync"
)

type Employee struct {
	Name       string
	Department string
	Salary     int
}

func main() {
	employees := []Employee{
		{Name: "Akhil", Department: "Dishwash", Salary: 13},
		{Name: "Dinesh", Department: "Dishwash", Salary: 13},
		{Name: "Rachana", Department: "Euro", Salary: 9},
		{Name: "Likitha", Department: "Salad", Salary: 9},
		{Name: "Sanjana", Department: "Grill", Salary: 9},
		{Name: "Sriven", Department: "Floor", Salary: 9},
	}

	// Here we are showcasing concurrency using Goroutines and Channels
	// Use a WaitGroup to wait for all Goroutines to finish
	var wg sync.WaitGroup

	// This is for creating a Channel of integers to receive salary values
	salaries := make(chan int)
	// Here we are showcasing exception handling using defer
	for _, employee := range employees {
		wg.Add(1)
		go func(emp Employee) {
			defer wg.Done()

			// If the employee's department is Dishwash, panic with a custom error message
			if emp.Department == "Dishwash" {
				panic(fmt.Sprintf("%s is in Dishwash department, cannot calculate salary", emp.Name))
			}
			salaries <- emp.Salary
		}(employee)
	}

	// Channel close
	go func() {
		wg.Wait()
		close(salaries)
	}()

	totalSalary := 0
	for salary := range salaries {
		totalSalary += salary
	}
	fmt.Printf("The total salary of all employees is %d dollars per hour\n", totalSalary)
}
