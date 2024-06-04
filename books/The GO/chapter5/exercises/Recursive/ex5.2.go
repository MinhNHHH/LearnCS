package exercises

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func FindLinkSs(filePath string) {
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
	docmentTree := map[string]int{}

	traverseLinks(docmentTree, doc)
	fmt.Println(docmentTree)
}

func traverseLinks(docmentTree map[string]int, n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode {
		tagName := n.Data
		docmentTree[tagName]++
	}

	// Recursively traverse the first child and then the next sibling
	traverseLinks(docmentTree, n.FirstChild)
	traverseLinks(docmentTree, n.NextSibling)
}
