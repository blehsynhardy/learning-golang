package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Print("=================== while loop style========== \n")
	//=============== while loop style =========//

	k := 3
	for k > 0 {
		fmt.Println(k)
		k--
	}

	fmt.Println("=================== infinite loop stye ============ \n")

	counter := 0
	for {
		fmt.Println(counter)
		counter++
		if counter >= 5 {
			break
		}
	}

	fmt.Print("=================== Array style ================ \n")
	items := [3]string{"A", "B", "C"}
	for index, item := range items {
		fmt.Println(index, item)
	}

}
