package exercises

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"golang.org/x/net/html"
)

// Extract makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML document.
func Extract87(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	// Ensure the directory exists
	err = os.MkdirAll(filepath.Dir("tmp/"), 0755)
	if err != nil {
		return nil, err
	}

	file, errCreate := os.Create(fmt.Sprintf("tmp/%s.html", sanitizeURL(url)))
	if errCreate != nil {
		return nil, fmt.Errorf("create file error %v", err)
	}
	err = html.Render(file, doc)
	if err != nil {
		return nil, err
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

//!-Extract

type LinksDepth struct {
	url   string
	depth int
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

// !+
func EX87() {
	worklist := make(chan []LinksDepth)  // lists of URLs, may have duplicates
	unseenLinks := make(chan LinksDepth) // de-duplicated URLs
	depthLimit := 10
	// Add command-line arguments to worklist.
	go func() {
		newWorkLists := []LinksDepth{}
		for _, url := range os.Args[1:] {
			newWorkLists = append(newWorkLists, LinksDepth{url: url, depth: 0})
		}
		worklist <- newWorkLists
	}()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for linksDepth := range unseenLinks {
				if linksDepth.depth < depthLimit {
					works := []LinksDepth{}
					foundLinks := crawl(linksDepth.url)
					for _, link := range foundLinks {
						works = append(works, LinksDepth{url: link, depth: linksDepth.depth + 1})
					}
					go func() { worklist <- works }()
				}
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link.url] {
				seen[link.url] = true
				unseenLinks <- link
			}
		}
	}
}

//!-
