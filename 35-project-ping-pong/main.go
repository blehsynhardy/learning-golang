package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, name string, c chan<- string, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that this goroutine is finished
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Shutting down %s...\n", name)
			return
		case c <- fmt.Sprintf("%s at %v", name, time.Now().Format("15:04:05")):
			time.Sleep(time.Second)
		}
	}
}

func main() {
	// 1. Setup Context with cancellation
	ctx, cancel := context.WithCancel(context.Background())
	
	pingChan := make(chan string)
	var wg sync.WaitGroup

	// 2. Start producers with a WaitGroup to track them
	wg.Add(2)
	go worker(ctx, "Ping", pingChan, &wg)
	go worker(ctx, "Pong", pingChan, &wg)

	// 3. Consumer logic
	go func() {
		timeout := time.After(5 * time.Second)
		for {
			select {
			case <-timeout:
				fmt.Println("--- Timeout reached, initiating shutdown ---")
				cancel() // Signal goroutines to stop
				return
			case msg := <-pingChan:
				fmt.Println(msg)
			}
		}
	}()

	// 4. Wait for workers to exit, then close the channel safely
	wg.Wait()
	close(pingChan)
	fmt.Println("All workers stopped. Done.")
}