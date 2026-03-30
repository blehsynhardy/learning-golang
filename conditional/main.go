package main

import "fmt"

func main() {
	userAccess := map[string]bool{
		"admin": true,
		"user":  false,
	}

	if hasAccess, ok := userAccess["admin"]; ok && hasAccess {
		fmt.Println("Admin access granted")
	}

	if userAccess["admin"] {
		fmt.Println("Admin")
	}

	//switch

	day := "Sunday"

	switch day {
	case "Sunday", "Saturday":
		fmt.Println("Weekend for my girlfriend")

	case "Monday", "Tuesday":
		fmt.Println("Congested day")

	default:
		fmt.Println("Mid week")
	}
}
