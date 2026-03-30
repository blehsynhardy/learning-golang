package main

import "fmt"

// Payable is a composite interface. 
// It embeds fmt.Stringer (requiring a String() method) and adds CalculatePay().
// Any type implementing both becomes a 'Payable' type.
type Payable interface {
	fmt.Stringer
	CalculatePay() float64
}

// Different employee types representing various payment structures
type Employee struct {
	ID     int
	Name   string
	Salary float64
}

type HourlyEmployee struct {
	ID          int
	Name        string
	HourlyRate  float64
	HoursWorked float64
}

type CommissionEmployee struct {
	ID             int
	Name           string
	BaseSalary     float64
	CommissionRate float64
	SalesAmount    float64
}

// Implementation for Salaried Employee (Monthly calculation)
func (e Employee) CalculatePay() float64 {
	return e.Salary / 12 // Assuming bi-weekly pay
}

func (e Employee) String() string {
	return fmt.Sprintf("Employee ID: %d, Name: %s, Salary: %.2f", e.ID, e.Name, e.Salary)
}

// Implementation for Hourly Employee
func (h HourlyEmployee) CalculatePay() float64 {
	return h.HourlyRate * h.HoursWorked
}

func (h HourlyEmployee) String() string {
	return fmt.Sprintf("Hourly Employee ID: %d, Name: %s, Hourly Rate: %.2f, Hours Worked: %.2f", h.ID, h.Name, h.HourlyRate, h.HoursWorked)
}

// Implementation for Commission-based Employee
func (c CommissionEmployee) CalculatePay() float64 {
	return c.BaseSalary + (c.CommissionRate * c.SalesAmount)
}

func (c CommissionEmployee) String() string {
	return fmt.Sprintf("Commission Employee ID: %d, Name: %s, Base Salary: %.2f, Commission Rate: %.2f, Sales Amount: %.2f", c.ID, c.Name, c.BaseSalary, c.CommissionRate, c.SalesAmount)
}

// printEmployeeSummary is a generic function.
// [P fmt.Stringer] ensures the input type P has a String() method.
func printEmployeeSummary[P fmt.Stringer](Employee P) {
	fmt.Printf("Processing employee: %s\n", Employee)
}

// processPayroll demonstrates polymorphism.
// It accepts a slice of the Payable interface, allowing it to handle any employee type.
func processPayroll(employees []Payable) {

	fmt.Println("\n====PROCESSING PAYROLL===========")

	totalPayroll := 0.0
	for _, employee := range employees {
		printEmployeeSummary(employee)
		pay := employee.CalculatePay()
		fmt.Printf("Calculated pay for %s: %.2f\n", employee, pay)
		totalPayroll += pay
	}
	fmt.Printf("Total payroll cost: %.2f\n", totalPayroll)
	fmt.Println("==================================")
}

func main() {
	fmt.Println("Welcome to the Payroll System")

	// Create some employee instances
	employee1 := Employee{ID: 1, Name: "Alice", Salary: 60000}
	employee2 := HourlyEmployee{ID: 2, Name: "Bob", HourlyRate: 15.50, HoursWorked: 160}
	employee3 := CommissionEmployee{ID: 3, Name: "Charlie", BaseSalary: 30000, CommissionRate: 0.10, SalesAmount: 50000}

	// Create a slice of employees
	employees := []Payable{employee1, employee2, employee3}

	// Process the payroll
	processPayroll(employees)

}
