package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
)

func check(name string) string {
	resp, err := http.Get("http://api" + name)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	return string(body)

}

func worker(name string, wg *sync.WaitGroup, names chan string) {
	defer wg.Done()
	var a = check(name)
	names <- a
}

func main() {

	names := make(chan string)
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker("www"+strconv.Itoa(i), &wg, names)
	}
	// Read from the channel until it is closed
	done := make(chan struct{})
	wg.Add(5)
	go func() {
		for x := range names {
			fmt.Println(x)
		}
		// Signal that println is completed
		close(done)
	}()

	// Wait for goroutines to end
	wg.Wait()
	// Close the channel to terminate the reader goroutine
	close(names)
	// Wait until println completes
	<-done
	fmt.Println(<-names)

}
