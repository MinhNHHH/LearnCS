package examples

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	// Recursively traverse the first child and then the next sibling
	links = visit(links, n.FirstChild)
	links = visit(links, n.NextSibling)
	return links
}

func FindLinkS1(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	doc, err := html.Parse(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println("asdasdsa", link)
	}
}
