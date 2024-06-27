package main

import (
	"fmt"
	"time"
)

func main() {
	wait := make(chan struct{}, 1)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Goroutine done")
		wait <- struct{}{} // Send a signal that the goroutine is done
	}()

	<-wait // Wait for the signal from the goroutine
	fmt.Println("Main function done")
}
