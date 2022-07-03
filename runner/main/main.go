// This sample program demostrates how to use a channel to
// monitor the amount of time the program is running and terminate
// the program if it runs too long.
package main

import (
	"log"
	"os"
	"time"

	"github.com/SiweiWang/go-in-action/runner"
)

const timeout = 5 * time.Second

func main() {
	log.Println("Starting work.")

	r := runner.New(timeout)

	r.Add(createTask(), createTask(), createTask(), createTask(), createTask())

	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout.")
			os.Exit(1)

		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt.")
			os.Exit(1)
		}
	}

	log.Println("Process ended.")
}

// createTask returns an example task that sleeps for the specified
// number of seconds based on the id.

func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)
		log.Printf("time to sleep %d seconda.", time.Duration(id))
		time.Sleep(time.Duration(id) * time.Second)
	}
}
