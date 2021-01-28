package search

import (
	"log"
)

// A map of registerd matchers for searching
var matchers = make(map[string]Matcher)

// Run performs the search logid.
func Run(searchTerm string) {
	// Retrieve the list of feeds to search through.
	_, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	// Create an unbuffered channel to receive match results to display.

}

func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatal(feedType, "Matcher already registered")

	}
}
