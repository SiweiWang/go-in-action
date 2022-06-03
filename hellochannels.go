package main

import (
	"fmt"
	"sync"
)

// waitgroup is used to wait for all the goroutines launched here to finish
var wg sync.WaitGroup

// simple print function
func printer(ch chan int) {
	for i := range ch {
		fmt.Printf("start go routine: %d \n", i)
	}

	// decrease the waitgroup counter by vaule of 1
	wg.Done()
}

func main() {
	c := make(chan int)

	// Define go routine
	go printer(c)

	// The waitgroup act as a counter holding the number of fucntion/go routines to wait for.
	// If the counter becomes 0, the wait group releasae the goroutine.
	wg.Add(1)

	for i := 1; i <= 20; i++ {
		c <- i
	}

	// close the channel c
	close(c)

	// The wait method blocks the execution until Waitgroup counter becomes 0.
	wg.Wait()
}
