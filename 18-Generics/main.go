package main


type number interface {
	int | float64
}
// sum is a generic function. 
// [T int | float64] defines a type parameter 'T' with a type constraint.
// This means the function can accept either a slice of integers or a slice of float64s.
// Generics allow us to write reusable code that works with multiple types while 
// maintaining type safety at compile time, avoiding the need for interface{} and type assertions.
func sum[T number](numbers ...T) T {
	// Initialize total with the zero value of type T
	total := T(0)
	for _, number := range numbers {
		total += number
	}
	return total
}

func main () {
	// Example with integers
	// The compiler infers that T is int
	println("Sum of ints:", sum(1, 2, 3, 4, 5))

	// Example with float64
	// The compiler infers that T is float64
	println("Sum of floats:", sum(1.5, 2.3, 5.8))
}
