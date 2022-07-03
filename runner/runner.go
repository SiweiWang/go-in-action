// Package runner manages the running and lifetime of a process
package runner

import (
	"errors"
	"os"
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
