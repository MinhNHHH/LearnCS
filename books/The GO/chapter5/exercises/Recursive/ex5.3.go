package exercises

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func PrintContent(filePath string) {
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

	printTextNodes(doc)
}

func printTextNodes(n *html.Node) {
	if n == nil {
		return
	}

	// If the node is a text node, print its content
	if n.Type == html.TextNode {
		fmt.Println(n.Data)
	}

	// If the node is a script or style element, do not traverse its children
	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		return
	}

	// Recursively traverse the first child and then the next sibling
	printTextNodes(n.FirstChild)
	printTextNodes(n.NextSibling)
}
