package main

import (
	"fmt"
	"sort"
)

func sortString(s string) string {
	// Convert string to a slice of runes to handle any Unicode characters.
	r := []rune(s)

	// Sort the slice of runes.
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})

	// Convert the sorted slice of runes back to a string.
	return string(r)
}

func canConstruct(ransomNote string, magazine string) bool {
	hashMagazine := map[rune]int{}
	hashRansom := map[rune]int{}

	for _, char := range magazine {
		hashMagazine[char]++
	}

	for _, char := range ransomNote {
		hashRansom[char]++
	}

	for key, value := range hashRansom {
		if hashMagazine[key] < value {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(canConstruct("aa", "aab"))
	fmt.Println(canConstruct("bg", "efjbdfbdgfjhhaiigfhbaejahgfbbgbjagbddfgdiaigdadhcfcj"))
	fmt.Println(canConstruct("abb", "ab"))
}
