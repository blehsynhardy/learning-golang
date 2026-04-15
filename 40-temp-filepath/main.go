package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	tempFile, err := os.CreateTemp("", "logs-*.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		fmt.Println("Removing file", tempFile.Name())
		if err := os.Remove(tempFile.Name()); err != nil {
			log.Fatal(err)
		}
	}()

	_, err = tempFile.Write([]byte("Hello world\n"))
	if err != nil {
		log.Fatal(err)
	}

	tempFile.Close()

	//temp direcotry

	tempDir, err := os.MkdirTemp("", "app-log-*")

	if err != nil {
		log.Fatal()
	}

	defer func() {
		fmt.Println("Removing directories")
		if err := os.RemoveAll(tempDir); err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Println("Temp dir created", tempDir)

}
