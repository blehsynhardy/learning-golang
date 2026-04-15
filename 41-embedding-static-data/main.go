package main

import (
	"embed"
	"fmt"
	"log"
)

//go:embed public

var public embed.FS

func main() {

	data, err := public.ReadFile("public/lead.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}

// import (
// 	_ "embed"
// 	"fmt"
// )

// //go:embed hello.txt


// var data string

// func main() {

// 	fmt.Println(data)

// }