package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

//base64 encoding and decoding

func main() {



	var data string = "weclome to the wonderful world of GO";


	encoded := base64.StdEncoding.EncodeToString([]byte(data))

	fmt.Println(encoded)	


	encodedString := "d2VjbG9tZSB0byB0aGUgd29uZGVyZnVsIHdvcmxkIG9mIEdP"

	decodedStr, err := base64.StdEncoding.DecodeString(encodedString)

	if err != nil {
		log.Fatal(err)
	}

	if string(decodedStr) != data {
		log.Fatal("encoding failed")
	}

	fmt.Println(string(decodedStr))



}