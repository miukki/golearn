package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func run(ctx context.Context) {
	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		for {
			select {
			case <-gCtx.Done():
				fmt.Println("Break the loop")
				return nil
			case <-time.After(1 * time.Second):
				fmt.Println("Hello in a loop")
			}
		}
	})

	g.Go(func() error {
		for {
			select {
			case <-gCtx.Done():
				fmt.Println("Break the loop")
				return nil
			case <-time.After(1 * time.Second):
				fmt.Println("Ciao in a loop")
			}
		}
	})

	err := g.Wait()
	if err != nil {
		fmt.Println("Error group: ", err)
	}
	fmt.Println("Main done")
}

func main() {
	// Create a context that is canceled after 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	run(ctx)
}
