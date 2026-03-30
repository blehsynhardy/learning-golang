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

//if we want to modify the struct field we need to use pointer receiver otherwise it will not update the struct field value and if we want to read the struct field value we can use value receiver or pointer receiver both will work fine but it is a good practice to use pointer receiver for methods that modify the struct fields and value receiver for methods that only read the struct fields.


// FullName is a method with a value receiver that concatenates the first and last name of an Employee.
func (e *Employee) FullName() string {
	return e.Firstname + " " + e.Lastname
}

// UpdateDepartment is a method with a pointer receiver that updates the department of an Employee.
func (e *Employee) UpdateDepartment(newDepartment string) {
	e.Department = newDepartment
}

func (e *Employee) Deactivate() {
	e.IsActive = false
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

	fmt.Println(jane.FullName())
	fmt.Println("Before update:", jane.Department)
	jane.UpdateDepartment("Marketing")
	fmt.Println("After update:", jane.Department)
	defer fmt.Println("Deactivating employee...")
}