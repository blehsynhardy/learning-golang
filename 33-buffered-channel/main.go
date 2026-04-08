package main

import (
	"fmt"
)

type Message struct {
	Text string
}

func main() {
	//unbuffered channels are channels that do not have a buffer, meaning that they can only hold one value at a time. When a value is sent to an unbuffered channel, the sender will block until the receiver is ready to receive the value. This can lead to deadlocks if the sender and receiver are not properly synchronized.

	//buffered channels are channels that have a buffer, meaning that they can hold multiple values at a time. When a value is sent to a buffered channel, the sender will not block until the buffer is full. This allows for more flexibility in the communication between goroutines, as the sender and receiver do not need to be perfectly synchronized.

	message := make(chan string, 3)

	fmt.Println("========SENDING MESSAGE TO BUFFER CHANNEL")

	message <- "Hello, World! This is a message from the goroutine."

	message <- "This is another message from the goroutine."

	message <- "This is a third message from the goroutine."

	fmt.Println(<-message)
	fmt.Println(<-message)
	fmt.Println(<-message)
}
