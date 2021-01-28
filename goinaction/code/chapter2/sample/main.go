package main

import (
	"log"
	"os"
)

// init is called prior to main
func init() {
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program
func main() {
	// Perform the search for the specified term
}
