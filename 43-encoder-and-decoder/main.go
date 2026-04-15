package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type name struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Phone     string `json:"phone_number"`
	Email     string `json:"email"`
}

func main() {

	person := name{
		Firstname: "blessing",
		Lastname:  "Atomicity",
		Phone:     "123-1212",
		Email:     "blessingawodele@gmail.com",
	}

	//this is more flexible than marshall because u can save the json data to anywhere stdout, buffer etc
	// for marhsall u can only store in bytes

	enc := json.NewEncoder(os.Stdout)

	if err := enc.Encode(person); err != nil {
		log.Fatal(err)
	}

	var payload = `{"firstname":"blessing","lastname":"Atomicity","phone_number":"123-1212","email":"blessingawodele@gmail.com"}`

	var n name 

	dec := json.NewDecoder(strings.NewReader(payload))
	if err := dec.Decode(&n); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", n)

	


	//final thought for web stuff its best to use marshall and unmarhsall when it comes to using the output dimension then use newEncoder and newDecoder


}
