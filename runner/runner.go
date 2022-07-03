// Package runner manages the running and lifetime of a process
package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// Runner runs a set of tasks within a given timeout and can be shut
// down on an operating system interrupt
type Runner struct {
	// interrupt channel reports a signal from the
	// Operating system.
	interrupt chan os.Signal

	// complete channel reports that processing is done
	complete chan error

	// timeout reports that time has run out.
	timeout <-chan time.Time

	// tasks holds a set of function that are executed
	// synchronously in index order
	tasks []func(int)
}

// ErrTimeout = errors.New("received timeout")
var ErrTimeout = errors.New("received timeout")

// ErrInterrupt is returned when an event from the OS is received
var ErrInterrupt = errors.New("received interrupt")

// New returns a new ready to use Runner
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

//Add attaches tasks to Runner. A task is a function that takes an init ID
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Start runs all tasks and monitors channel events.
func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	select {
	// signaled when processing is done
	case err := <-r.complete:
		return err

	// singaled when we run out of time
	case <-r.timeout:
		return ErrTimeout
	}
}

// run executes each registered task.
func (r *Runner) run() error {
	for id, task := range r.tasks {
		// Check for an interrupt singal from the OS.
		if r.gotInterrupt() {
			return ErrInterrupt
		}

		//execute the registered task
		task(id)
	}
	return nil
}

// gotInterrupted verfies if the interrupt signal has been issued
func (r *Runner) gotInterrupt() bool {
	select {
	//signaled when a interrupt event is sent
	case <-r.interrupt:
		// Stop receiving any further singals.
		signal.Stop(r.interrupt)
		return true

	// Continue running as normal
	default:
		return false
	}
}
