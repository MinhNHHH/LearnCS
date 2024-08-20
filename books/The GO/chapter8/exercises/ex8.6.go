// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 243.

// Crawl3 crawls web links starting with the command-line arguments.
//
// This version uses bounded parallelism.
// For simplicity, it does not address the termination problem.
//
package exercises

import (
	"log"
	"os"
)

type LinkDepth struct {
	url   string
	depth int
}

func Crawl(url string) []string {
	list, err := Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

//!+
func Ex86() {
	worklist := make(chan []LinkDepth)  // lists of URLs, may have duplicates
	unseenLinks := make(chan LinkDepth) // de-duplicated URLs
	depthLimit := 10
	// Add command-line arguments to worklist.
	go func() {
		var initialLinks []LinkDepth
		for _, url := range os.Args[1:] {
			initialLinks = append(initialLinks, LinkDepth{url, 0})
		}
		worklist <- initialLinks
	}()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for linkDepth := range unseenLinks {
				if linkDepth.depth < depthLimit {
					foundLinks := Crawl(linkDepth.url)
					var newLinks []LinkDepth
					for _, link := range foundLinks {
						newLinks = append(newLinks, LinkDepth{link, linkDepth.depth + 1})
					}
					go func() { worklist <- newLinks }()
				}
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, linkDepth := range list {
			if !seen[linkDepth.url] {
				seen[linkDepth.url] = true
				unseenLinks <- linkDepth
			}
		}
	}
}

//!-
