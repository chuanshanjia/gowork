package main

import (
	"github.com/chuanshanjia/gowork/goinaction/code/chapter2/sample/search"
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
	search.Run("president")
}
