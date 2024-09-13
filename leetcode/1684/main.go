package main

import (
	"fmt"
)

func countConsistentStrings(allowed string, words []string) int {
	count := 0
	set := make([]bool, 26)

	for _, c := range allowed {
		set[c-'a'] = true
	}

	for _, word := range words {
		for _, char := range word {
			if !set[char-'a'] {
				count++
				break
			}
		}
	}
	return len(words) - count
}

func main() {
	//fmt.Println(countConsistentStrings("ab", []string{"a", "bd", "aaab", "baa", "badab"}))
	fmt.Println(countConsistentStrings("fstqyienx", []string{"n", "eeitfns", "eqqqsfs", "i", "feniqis", "lhoa", "yqyitei", "sqtn", "kug", "z", "neqqis"}))
}
