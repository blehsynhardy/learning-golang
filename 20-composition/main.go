package main

import "fmt"

//Composition is a design principle in which a class is composed of one or more objects from other classes, rather than inheriting from a parent class. In Go, we can achieve composition using struct embedding and interfaces. This allows us to create flexible and reusable code.

type Address struct {
	Street string
	City   string
	State  string
	Zip    string
}

type Customer struct {
	CustomerID     int
	Name           string
	Email          string
	BillingAddress Address
}

func (a Address) FullAddress() string {
	if a.Street == "" || a.City == "" || a.State == "" || a.Zip == "" {
		return "Address is incomplete"
	}
	return a.Street + ", " + a.City + ", " + a.State + " " + a.Zip
}

func (c Customer) ContactInfo() string {
	fmt.Printf("Customer ID :%d\n", c.CustomerID);
	fmt.Printf("Customer Name :%s\n", c.Name);
	fmt.Printf("Customer Email :%s\n", c.Email);
	fmt.Printf("Customer Billing Address :%s\n", c.BillingAddress.FullAddress()); // This will call the FullAddress method of the Address struct to get the complete address as a string.

	return fmt.Sprintf("Customer ID: %d, Name: %s, Email: %s, Billing Address: %s", c.CustomerID, c.Name, c.Email, c.BillingAddress.FullAddress())
}

func main() {
	// In this example, we will demonstrate composition in Go by creating a simple application that models a company with employees and customers. We will use interfaces to define common behavior and struct embedding to achieve composition.

	fmt.Println("========COMPOSITION==========")
	// Create an instance of Address
	customer1 := Customer{
		CustomerID: 1,
		Name:       "John Doe",
		Email:      "customer@gmail.com",
		BillingAddress: Address{
			Street: "123 Main St",
			City:   "Anytown",
			State:  "CA",
			Zip:    "12345",
		},
	}

	// Display the contact information of the customer
	fmt.Println(customer1.ContactInfo())

}
