package main

import "fmt"

func main() {
	// A pointer is a variable instead of storing a value it store an address of another variable
	age := 10

	fmt.Printf("Age is %d\n", &age)
}
