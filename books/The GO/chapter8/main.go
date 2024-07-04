package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create a context with cancellation
	ctx, cancel := context.WithCancel(context.Background())

	// Start a goroutine that listens for the cancellation signal
	go func() {
		select {
		case <-time.After(5 * time.Second):
			fmt.Println("completed work")
		case <-ctx.Done():
			fmt.Println("work canceled")
		}
	}()

	// Simulate some work in the main goroutine
	time.Sleep(6 * time.Second)

	// Cancel the context
	cancel()

	// Give the goroutine time to exit
	time.Sleep(1 * time.Second)
}
