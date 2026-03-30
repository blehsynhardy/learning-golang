package main

import "fmt"

func main() {
	fmt.Println("================== ARRAY ===")
	var productList [2]string

	productList[0] = "blessing"
	productList[1] = "tommy"
	fmt.Printf("%v\n", productList)

	primes := [3]int{2, 3, 5}

	for i := 0; i < len(primes); i++ {
		fmt.Printf("Prime = %v\n", primes[i])
	}

	fmt.Println("====================== SLICE ===============")
	name := []string{"John", "Doe", "Bob"}

	fmt.Printf("%v\n", name)
	makeName := make([]int, 3, 10)

	fmt.Printf("%v\n", makeName)

	fmt.Println("====================== MAP SAME AS OBJECT ===============")

	userInfo := map[string]int{
		"john": 10,
		"doe":  20,
		"bob":  30,
	}
	fmt.Printf("%v\n", userInfo)

	userInfo["john"] = 20
	fmt.Printf("%v\n", userInfo)

	checkJohn, ok := userInfo["john"]

	if ok {
		fmt.Printf("%v\n", checkJohn)
	}

	if _, ok := userInfo["bob"]; ok {
		fmt.Printf("%v\n", userInfo["bob"])
	}

	config := make(map[string]int)

	config["one"] = 1
}
