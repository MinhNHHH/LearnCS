package main

import (
	"fmt"
	"strings"
)

func isCircularSentence(sentence string) bool {
	newSentences := strings.Split(sentence, " ")
	if len(newSentences) == 1 {
		return newSentences[0][0] == newSentences[0][len(newSentences)-1]
	}
	n := len(newSentences)

	if newSentences[0][0] != newSentences[n-1][len(newSentences[n-1])-1] {
		return false
	}
	for i := 0; i < len(newSentences)-1; i++ {
		current := newSentences[i]
		next := newSentences[i+1]

		if current[len(current)-1] != next[0] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isCircularSentence("leetcode exercises sound delightful"))
}
