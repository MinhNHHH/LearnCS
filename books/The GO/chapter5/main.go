package main

import (
	"chapter5/exercises/functv"
	"fmt"
	"os"

	"golang.org/x/net/html"
)

const filePath = "mock-data/draft.html"

func main() {
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

	newPrettyPrinter := functv.NewPrettyPrinter()
	newPrettyPrinter.Print(doc)

	prettyHTML := newPrettyPrinter.String()
	fmt.Println(prettyHTML)

	// Write the pretty-printed HTML to a file
	outputFile := "output.html"
	err = os.WriteFile(outputFile, []byte(prettyHTML), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing file: %v\n", err)
		os.Exit(1)
	}
}
