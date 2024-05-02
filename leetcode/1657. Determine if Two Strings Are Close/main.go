// Two strings are considered close if you can attain one from the other using the following operations:

//     Operation 1: Swap any two existing characters.
//         For example, abcde -> aecdb
//     Operation 2: Transform every occurrence of one existing character into another existing character, and do the same with the other character.
//         For example, aacabb -> bbcbaa (all a's turn into b's, and all b's turn into a's)

// You can use the operations on either string as many times as necessary.

// Given two strings, word1 and word2, return true if word1 and word2 are close, and false otherwise.

// Example 1:

// Input: word1 = "abc", word2 = "bca"
// Output: true
// Explanation: You can attain word2 from word1 in 2 operations.
// Apply Operation 1: "abc" -> "acb"
// Apply Operation 1: "acb" -> "bca"

// Example 2:

// Input: word1 = "a", word2 = "aa"
// Output: false
// Explanation: It is impossible to attain word2 from word1, or vice versa, in any number of operations.

// Example 3:

// Input: word1 = "cabbba", word2 = "abbccc"
// Output: true
// Explanation: You can attain word2 from word1 in 3 operations.
// Apply Operation 1: "cabbba" -> "caabbb"
// Apply Operation 2: "caabbb" -> "baaccc"
// Apply Operation 2: "baaccc" -> "abbccc"

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(closeStrings("uau", "xax"))
}

func sortString(word string) string {
	s := []rune(word)
	sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })
	return string(s)
}
func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	hash1 := map[rune]int{}
	s1 := map[rune]bool{}
	hash2 := map[rune]int{}

	for _, char := range word1 {
		hash1[char]++
		s1[char] = true
	}

	for _, char := range word2 {
		hash2[char]++
		if _, ok := s1[char]; !ok {
			return false
		}
	}

	count1 := []int{}
	for _, value := range hash1 {
		count1 = append(count1, value)
	}
	sort.Ints(count1)

	count2 := []int{}
	for _, value := range hash2 {
		count2 = append(count2, value)
	}
	sort.Ints(count2)

	for index, value := range count1 {
		if value != count2[index] {
			return false
		}
	}
	return true
}
