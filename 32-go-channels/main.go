package main

import (
	"fmt"
	"time"
)

type Message struct {
	Text string
}

func main() {
	//explain channels in go and how to use them
	//Channels are a powerful feature in Go that allow for communication and synchronization between goroutines. They provide a way to send and receive values between different parts of a program, enabling concurrent programming. Channels can be used to coordinate the execution of goroutines, allowing them to communicate and share data safely without the need for locks or other synchronization mechanisms.

	message := make(chan string)

	data := make(chan Message)

	go func() {
		message <- "Hello, World! This is a message from the goroutine."
	}()

	go func() {
		data <- Message{Text: "This is a message struct from the goroutine in struct."}
	}()

	time.Sleep(1 * time.Second);

	fmt.Println("Waiting for message...")

	msg := <-message

	fmt.Println("Received message:", msg)
	fmt.Println("Waiting for data...")
	dataMsg := <-data
	fmt.Println("Received data:", dataMsg)
}