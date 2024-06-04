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

	contents := traverseLinkC([]string{}, doc)
	for _, content := range contents {
		fmt.Println(content)
	}
}

func traverseLinkC(contents []string, n *html.Node) []string {
	if n == nil {
		return contents
	}
	fmt.Println(n.Data == "script" || n.Data == "style")
	if n.Type == html.ElementNode && (n.Data != "script" || n.Data != "style") {
		for _, c := range n.Attr {
			contents = append(contents, c.Val)
		}
	}

	// Recursively traverse the first child and then the next sibling
	contents = traverseLinkC(contents, n.FirstChild)
	contents = traverseLinkC(contents, n.NextSibling)
	return contents
}
