package main

import (
	"fmt"
	"strings"
)

func main() {

	myAccount := "12345"

	// strings.Clone creates a fresh copy of the string
	newaccount := strings.Clone(myAccount)
	fmt.Println(myAccount, newaccount)

	// Changing string casing
	fmt.Println(strings.ToUpper(myAccount))
	fmt.Println(strings.ToLower(myAccount))

	// strings.TrimSpace removes leading and trailing whitespace
	s4 := "   Hello World   "
	fmt.Printf("'%s' (length: %d)\n", s4, len(s4))
	fmt.Printf("'%s' (length: %d)\n", strings.TrimSpace(s4), len(strings.TrimSpace(s4)))

	// strings.HasPrefix checks if a string starts with a specific substring
	fmt.Println(strings.HasPrefix(myAccount, "123"))

	// Iterating through a slice to check suffixes
	s9 := []string{"@gmail.com", "@yahoo.com", "@outlook.com"}
	testEmail := "blessingawodele@gmail.com"

	for _, domain := range s9 {
		// strings.HasSuffix checks if a string ends with a specific substring
		if strings.HasSuffix(testEmail, domain) {
			fmt.Printf("The email '%s' ends with '%s'\n", testEmail, domain)
		} 
	}

	// strings.Replace replaces a specific number of occurrences of a substring
	fmt.Println(strings.Replace("test@gmail.com", "test", "john", 1))

	// strings.Split breaks a string into a slice based on a separator
	parts := strings.Split("scotthruska@gmail.com", "@")
	username, domain := parts[0], parts[1]
	fmt.Printf("Username: %s, Domain: %s\n", username, domain)

	// strings.Join concatenates elements of a slice into a single string with a separator
	part3 := strings.Join([]string{"scott", "hruska"}, "-")

	// Manual concatenation using the + operator
	part6 := "scott" + "-" + "hruska"

	fmt.Println(part3)
	fmt.Println(part6)

	// strings.Fields splits a string around each instance of one or more consecutive white space characters
	part4 := strings.Fields("jane example")
	fmt.Println(part4)
}
