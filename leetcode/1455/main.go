package main

import (
	"fmt"
	"strings"
)

func isPrefixOfWord(sentence string, searchWord string) int {
	splitSenteces := strings.Split(sentence, " ")
	for index, word := range splitSenteces {
		if strings.HasPrefix(word, searchWord) {
			return index + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(isPrefixOfWord("hellohello hellohellohello", "ell"))
}
