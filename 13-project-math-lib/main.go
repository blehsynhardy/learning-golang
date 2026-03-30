package main

import (
	"fmt"
	"strings"
)

// MathError defines a custom structure to hold detailed information about mathematical errors.
type MathError struct {
	Op      string
	InputA  int
	InputB  int
	Message string
}

const (
	division = "Division"
	ErrDivideByZero = "cannot divide by zero"
)

// Error implements the error interface for the MathError pointer receiver.
// It formats the error message to include the operation type and the inputs involved.
func (e *MathError) Error() string {
	var inputs []string

	if e.Op == division {
		// Collect input values to provide context in the error message
		inputs = append(inputs, fmt.Sprintf("InputA: %d", e.InputA))
		inputs = append(inputs, fmt.Sprintf("InputB: %d", e.InputB))
	}

	return fmt.Sprintf("Math error in %s (%s): %s", e.Op, strings.Join(inputs, ","), e.Message)

}

// sum takes a variadic number of integers and returns their total.
func sum(numbers ...int) int {
	// defer ensures this message prints after the function logic completes
	defer fmt.Println("Sum calculated finished")
	
	total := 0
	for _, num := range numbers {
		total += num
	}

	return total

}

// SafeDivision performs integer division but returns a custom MathError if the divisor is zero.
func SafeDivision(a, b int) (int, error) {

	// Check for division by zero to prevent a runtime panic
	if b == 0 {
		return 0, &MathError{
			Op:      division,
			InputA:  a,
			InputB:  b,
			Message: ErrDivideByZero,
		}
	}
	return a / b, nil	
}

func main() {
	// Example of using the variadic sum function
	fmt.Println("Sum of 1, 2, 3:", sum(1, 2, 3))

	// Example of handling a custom error from SafeDivision
	result, err := SafeDivision(10, 0)
	if err != nil {
		fmt.Println("Error:", err) // This calls the .Error() method automatically
		return
	}
	
	fmt.Println("Result of division:", result)
}
