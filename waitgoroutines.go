package main

import (
	"context"
	"fmt"
	"time"
)

func run(ctx context.Context) {
	wait := make(chan struct{}, 2) // Changed to 2 to wait for both goroutines

	go func() {
		defer func() {
			wait <- struct{}{}
		}()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Break the loop")
				return // Deferred function will execute here, Use return to exit the goroutine
			case <-time.After(1 * time.Second):
				fmt.Println("Hello in a loop")
			}
		}
	}()

	go func() {
		defer func() {
			wait <- struct{}{}
		}()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Break the loop")
				return // Deferred function will execute here, Use return to exit the goroutine
			case <-time.After(1 * time.Second):
				fmt.Println("Ciao in a loop")
			}
		}
	}()

	// wait for both goroutines to finish
	<-wait
	<-wait

	fmt.Println("Main done")
}

func main() {
	// Create a context that is canceled after 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	run(ctx)
}
