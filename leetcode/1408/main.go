package main

import (
	"fmt"
	"strings"
)

// Give an array of string words ["mass", "as", "hero", "superhero"]
// return an array of string that is an substring in another word => ["as", "hero"]

func stringMatching(words []string) []string {
	res := []string{}

	for _, word := range words {
		for _, compareWord := range words {
			if word != compareWord && strings.Contains(compareWord, word) {
				res = append(res, word)
				break
			}
		}
	}
	return res
}

func main() {
	fmt.Println(stringMatching([]string{"leetcoder", "leetcode", "od", "hamlet", "am"}))
}
