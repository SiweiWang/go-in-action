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
	waitGroup.Add(len(feeds))
	
	// Launch a goroutine for each feed to find the results
	for _, feed := range feeds{
		// Retrieve a matcher for the search
		if !exists {
			matcher = matchers["default"]	
		}
		
		// Launch the goroutine to perform the search
		go func(matcher Matcher, feed *Feed) {
			Match(match, feed, searchTerm, results)
			waitGroup.Done()	
		}(matcher, feed)
		
	}
	
	// Launch a goroutine to monitor when all the work is done
	go func() {
		// Wait for everything to be processed
		waitGroup.wait()
		
		// Close the channel to signal to the Display
		// function that we can exit the program
		close(results)
	}()
	
	// Start displaying results as they are avaiable and
	// return after the final result is displayed
	Display(results)
}

// Register is called to register a matcher for use by the program
func Register(feedType string, matcher Matcher) {
	if _, exists := matcher[feedType]; exists{
		log.Fatalln(feedType, "Matcher already registered")	
	}
	
	log.Println("Register", feedtype, "matcher")
	matchers[feedType] = matcher
}
