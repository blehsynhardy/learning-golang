package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {


	//using the writefile module to create file

	var filePath string = "text.txt"

	data := "lets write this to a file new file"

	err := os.WriteFile(filePath, []byte(data), 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println("done");

	content, err := os.ReadFile(filePath);

	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))



		//using the OS module to create file 
	// fmt.Print("==========FILE 2=========")

	// file2,err := os.Create("fle-created-with-os-m.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer file2.Close()

	// _, err = file2.WriteString("welcome from node js dev");
	// if err != nil {
	// 	panic(err)
	// }

	//newFilePath := "./file-created-with-os-m.txt"


	fileOpen, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	defer fileOpen.Close()

	scanner := bufio.NewScanner(fileOpen)
	lineNum := 1

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("%d: %s\n", lineNum, line)
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		if err != io.EOF {
			log.Fatal(err)
		}

	}



}
