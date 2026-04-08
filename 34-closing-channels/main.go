package main

import (
	"fmt"
	"sync"
)

type Message struct {
	Text string
}

//see the function at the bottom first, then the main function

func main() {
	//we want to use wait groups to wait for the worker to finish, but we can also use channels to signal when the worker is done. In this example, we will use a channel to signal when the worker is done. We will create a channel called done, and we will send a value to the done channel when the worker is done. The main function will wait for a value to be sent to the done channel before exiting.

	jobs := make(chan string, 3)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			job, more := <-jobs
			if more {
				fmt.Println("Received job:", job)
			} else {
				fmt.Println("No more jobs, worker is done.")
				return
			}
		}
	}(&wg)

	for i := 1; i <= 3; i++ {
		jobs <- fmt.Sprintf("Job %d", i)
	}
	close(jobs)

	wg.Wait()
	fmt.Println("Sent all jobs, waiting for worker to finish...")

}

func ChannelClosingExample() {
	//example of closing a channel in go
	//closing channels in go is a way to signal that no more values will be sent on a channel. This is done using the built-in close function. When a channel is closed, any subsequent sends on that channel will panic, and any receives will return the zero value of the channel's type. Closing a channel is useful for signaling to goroutines that they should stop waiting for values on that channel, and it can also be used to indicate that a stream of data has ended.

	jobs := make(chan string, 3)
	done := make(chan bool)

	go func() {
		for {
			job, more := <-jobs
			if more {
				fmt.Println("Received job:", job)
			} else {
				fmt.Println("No more jobs, worker is done.")
				done <- true
				return
			}
		}
	}()

	for i := 1; i <= 3; i++ {
		jobs <- fmt.Sprintf("Job %d", i)
	}
	close(jobs)
	fmt.Println("Sent all jobs, waiting for worker to finish...")

	<-done

}
