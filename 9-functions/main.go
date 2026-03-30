package main

import (
	"errors"
	"fmt"
	"strings"
)

func sum(numbers ...int) int {
	total := 0

	for _, n := range numbers {
		total += n
	}
	return total
}

//func divide(a int, b int) int {
//	return a / b
//}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Cannot be divided by zero")
	}
	return a / b, nil
}

func splitNames(fullname string) (firstname, lastname string) {
	parts := strings.Split(fullname, " ")
	firstname = parts[0]
	lastname = parts[1]

	return 
}

func main() {

	//sum(1, 2, 3, 4, 5)
	value, err := divide(10, 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(value)

	firstname, _ := splitNames("Awodele Blessing")

	fmt.Println(firstname)
}
