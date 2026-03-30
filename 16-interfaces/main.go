package main

// Employee represents a staff member with personal and professional details.
type Employee struct {
	ID   int
	Name string
}

// Customer represents a client with personal details.
type Customer struct {
	ID   int
	Name string
}

// Person is an interface that defines a contract for any type that has a getName method.
// This allows for polymorphism where different types can be treated as a 'Person'.
type Person interface {	
	getName() string
}

// getName implements the Person interface for the Employee struct.
func (e Employee) getName() string {
	return e.Name
}

// getName implements the Person interface for the Customer struct.
func (c Customer) getName() string {
	return c.Name
}

// displayPerson accepts any type that satisfies the Person interface.
// This demonstrates abstraction; the function doesn't care if it's an Employee or Customer.
func displayPerson(p Person) {
	println("Name:", p.getName())
}	

func main() {
	// Initialize an Employee instance
	joel := Employee{ID: 1, Name: "Joel"}
	
	// Initialize a Customer instance
	joe := Customer{ID: 2, Name: "Alice"}

	// Both Employee and Customer can be passed to displayPerson 
	// because they both implement the getName() method required by the Person interface.
	displayPerson(joel)
	displayPerson(joe)
}
