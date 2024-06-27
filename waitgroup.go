package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func run(ctx context.Context) {
	var wg sync.WaitGroup

	// Start the first goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Break the loop")
				return
			case <-time.After(1 * time.Second):
				fmt.Println("Hello in a loop")
			}
		}
	}()

	// Start the second goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Break the loop")
				return
			case <-time.After(1 * time.Second):
				fmt.Println("Ciao in a loop")
			}
		}
	}()

	// Wait for both goroutines to finish
	wg.Wait()
	fmt.Println("Main done")
}

func main() {
	// Create a context that is canceled after 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	run(ctx)
}
