package main

import "fmt"

// Address defines a basic physical location structure
type Address struct {
	Street string
	City   string
	State  string
	Zip    string
}

// ContactInfo represents personal contact details and links a billing address
type ContactInfo struct {
	CustomerID     int
	Name           string
	Email          string
	BillingAddress Address
}

// Company demonstrates struct embedding (composition)
// It embeds Address and ContactInfo to "inherit" their fields and methods
type Company struct {
	Address     // Anonymous field: fields/methods of Address are promoted to Company
	ContactInfo // Anonymous field: fields/methods of ContactInfo are promoted to Company
	// Email 	 string // Embedding Email directly in Company for easy access it shadwows the Email field from ContactInfo
	ID           int
	CompanyName  string
	BusinessType string
}

// FullAddress returns a formatted string of the Address fields
func (a Address) FullAddress() string {
	if a.Street == "" && a.City == "" && a.State == "" && a.Zip == "" {
		return "No address provided"
	}
	return a.Street + ", " + a.City + ", " + a.State + " " + a.Zip
}

// ContactDetails returns a formatted string of the ContactInfo fields
func (c ContactInfo) ContactDetails() string {
	return "Name: " + c.Name + "\nEmail: " + c.Email + "\nBilling Address: " + c.BillingAddress.FullAddress()
}

// CompanyDetails prints and returns company information, demonstrating access to promoted fields
func (c Company) CompanyDetails() string {
	fmt.Printf("Company Name: %s\n", c.CompanyName)
	fmt.Printf("Address : %s\n", c.FullAddress()) // Accessing promoted method from Address
	fmt.Printf("Street : %s\n", c.Street)           // Accessing promoted field from Address
	fmt.Printf("City : %s\n", c.City)               // Accessing promoted field from Address
	fmt.Printf("Email : %s\n", c.Email)             // Accessing promoted field from ContactInfo
	fmt.Printf("Business Type : %s\n", c.BusinessType)

	return "ye"
}

func main() {

	fmt.Println("========STRUC EMBEDDING============")

	// Initialize ContactInfo
	johnContact := ContactInfo{
		CustomerID: 1,
		Name:       "John Doe",
		Email:      "tommy@gmail.com",
		BillingAddress: Address{
			Street: "456 Elm St",
			City:   "Othertown",
			State:  "NY",
			Zip:    "54321",
		},
	}

	// Initialize Company using the previously created ContactInfo and a new Address
	johnCompany := Company{
		Address: Address{
			Street: "123 Main St",
			City:   "Anytown",
			State:  "CA",
			Zip:    "12345",
		},
		ContactInfo: johnContact,
		ID:           1,
		CompanyName:  "John Doe's Company",
		BusinessType: "Software Development",
	}

	// Display company details (uses promoted fields/methods)
	fmt.Println(johnCompany.CompanyDetails())
	fmt.Println("=============================================")
	
	// Display contact details (uses promoted method from ContactInfo)
	fmt.Println(johnCompany.ContactDetails())
}
