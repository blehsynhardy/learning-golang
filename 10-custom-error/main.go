package main

import (
	"errors"
)

/*
instead of hardcoding the error message in the function, we can create a variable to hold the error message and return that variable instead. This way, we can reuse the error message in multiple places and also make it easier to change the error message in the future if needed.
*/

var ErrDivideByZero = errors.New("Cannot be divided by zero")
var ErrorNumberTooLarge = errors.New("Number is too large")

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}

	if a > 1000 {
		return 0, ErrorNumberTooLarge
	}
	return a / b, nil
}

func main() {

	value, err := divide(100008, 1)
	if err != nil {
		if errors.Is(err, ErrDivideByZero) {
			println("Cannot divide by zero")
		}
		if errors.Is(err, ErrorNumberTooLarge) {
			println("Number is too large")
		}
		return
	}
	println(value)

}
