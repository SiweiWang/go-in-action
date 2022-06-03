package search

import (
	"log"
	"sync"
)

// A map of registered matchers for searching.
var matchers = make(map[string]Matcher)

// Run performs the search logic
func Run(seachTerm string) {
	//Retrieve the list of feeds to search through
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	// Create an unbuffered channel to receive match results to display
	results := make(chan *Result)

	// Setup a wait group so we can process all the feeds
	var waitGroup sync.WaitGroup

	// Setup the number of goroutines we need

}
