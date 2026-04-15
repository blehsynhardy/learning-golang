package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type name struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	State     string `json:"state"`
	Password  string `json:"password,omitempty"`
	Profile   profile `json:"profile"`
}

type profile struct {
	URL string `json:"url"`
}

//Marshall and Unmarshall means converting struct to a json or xml data or changing json data to struct data

func main() {

	person := name{
		Firstname: "blessing",
		Lastname:  "Atomicity",
		State:     "Lagos",
		Profile: profile{
			URL: "httsp://1.com",
		},
	}

	personJson, err := json.MarshalIndent(person, " ", " ") //use Marshall if u are not reading

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(personJson))

	//unMarshall

	var payload = ` {
	"firstname": "blessing",
	"lastname": "Atomicity",
	"state": "Lagos",
	"profile" : {
		"url" : "httsp://1.com"
		}
 	}`

	var n name

	err = json.Unmarshal([]byte(payload), &n)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", n)

}
