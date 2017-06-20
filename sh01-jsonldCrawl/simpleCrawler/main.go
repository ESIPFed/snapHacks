package main

import (
	"fmt"
)

// A simple crawler to go through a given web site (single domain) and starting at
// a JSON-LD document, walk through the tree of documents.  Driven by either hydra or
// by JSON-LD framing.  We wil try both
func main() {
	fmt.Println("Simple crawler")

	// Take a seed domain
	// read int he JSON-LD  (validate it, then frame it against a given frame...  place results into a struct)
	// Load the URLs discovered to crawl to a boltdb system (where I can check if they have been already crawled)
	// Store the JSON-LD to a graph as triples.  (also store the original JSON-LD to a bolt table.)
	// In the end we have the triples of the site...
}
