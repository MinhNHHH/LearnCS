package main

import (
	"fmt"
	"strings"
)

func wordSubsets(words1 []string, words2 []string) []string {
	res := []string{}
	hashMap := map[rune]bool{}

	for _, word := range words2 {
		for _, char := range word {
			hashMap[char] = true
		}
	}

	for _, word := range words1 {
		for key := range hashMap {
			if strings.Contains(word, string(key)) {
			} else {
				break
			}
		}
	}
	return res
}

func main() {
	fmt.Println(wordSubsets([]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"lo", "eo"}))
	fmt.Println(wordSubsets([]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"oo", "e"}))

}
