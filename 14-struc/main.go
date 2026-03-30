package main

import (
	"fmt"
	"time"
)

// Employee represents a staff member with personal and professional details.
type Employee struct {
	Firstname  string
	Lastname   string
	Age        int
	Department string
	Salary     float64
	IsActive   bool
	Time       time.Time
}

// NewEmployee is a constructor function that initializes a new Employee struct with the current timestamp.
func NewEmployee(firstname, lastname string, age int, department string, salary float64, isActive bool) Employee {
	return Employee{
		Firstname:  firstname,
		Lastname:   lastname,
		Age:        age,
		Department: department,
		Salary:     salary,
		IsActive:   isActive,
		Time:       time.Now(),
	}
}


func main() {

	// Initializing a struct using a composite literal
	jane := Employee{
		Firstname: "Jane",
		Lastname: "Doe",
		Age: 30,
		Department: "Engineering",
		Salary: 75000.00,
		IsActive: true,
		Time: time.Now(),
	}

	// Initializing a struct using the constructor function
	john := NewEmployee("John", "Smith", 28, "Marketing", 65000.00, true)
	fmt.Printf("Employee: %+v\n", john)

	// Modifying struct fields directly
	// janePtr := &jane 
	// janePtr.Firstname = "BLessing";
	jane.Firstname = "Atomicity";

	fmt.Printf("Employee: %+v\n", jane)
}