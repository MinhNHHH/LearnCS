package main

import "fmt"

func countInRange(rangeMin, rangeMax int) int {
	count := 0
	for i := rangeMin; i <= rangeMax; i++ {
		count++
	}
	return count
}

func vowelStrings(words []string, queries [][]int) []int {
	countVowel := make([]int, len(queries))
	vowelString := map[string]bool{
		"a": true,
		"e": true,
		"i": true,
		"o": true,
		"u": true,
	}
	temp := 0
	prefixSum := make([]int, len(words))
	for index, word := range words {
		if vowelString[string(word[0])] && vowelString[string(word[len(word)-1])] {
			temp++
		}
		prefixSum[index] = temp
	}

	for i, query := range queries {
		left, right := query[0], query[1]
		if left >= 1 {
			countVowel[i] = prefixSum[right] - prefixSum[left-1]
		} else {
			countVowel[i] = prefixSum[right]
		}
	}

	return countVowel
}

func main() {
	fmt.Println(vowelStrings([]string{"aba", "bcb", "ece", "aa", "e"}, [][]int{{0, 2}, {1, 4}, {1, 1}}))
}
