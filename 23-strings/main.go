package main

import (
	"fmt"
	"strings"
)

func main() {

	myAccount := "12345"

	//fmt.Println(myAccount);

	newaccount := strings.Clone(myAccount)

	fmt.Println(myAccount, newaccount)
}
