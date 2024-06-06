package multiple

import (
	"fmt"
	"net/http"
	"strings"
	"unicode"

	"golang.org/x/net/html"
)

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, err
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return 0, 0, fmt.Errorf("parsing HTML: %s", err)
	}
	words, images = countWordsAndImages(doc)

	return words, images, nil
}

func countWordsAndImages(n *html.Node) (words, images int) {
	// countWords := map[string]int{}
	// countImages := 0
	countWords, countImages := traverseLinkz(n, 0, 0)
	fmt.Println(countWords, countImages)
	return countWords, countImages
}

func isAlpha(str string) bool {
	for _, char := range str {
		if !unicode.IsLetter(char) {
			return false
		}
	}
	return true
}

func traverseLinkz(n *html.Node, countWords, countImages int) (int, int) {
	if n == nil {
		return countWords, countImages
	}

	if n.Type == html.ElementNode && n.Data != "img" {
		countImages++
	} else if n.Type == html.TextNode {
		textList := strings.Split(n.Data, " ")
		for _, text := range textList {
			if isAlpha(text) {
				countWords++
			}
		}
	}

	countWords, countImages = traverseLinkz(n.FirstChild, countWords, countImages)
	countWords, countImages = traverseLinkz(n.NextSibling, countWords, countImages)
	return countWords, countImages
}
