package main

import "fmt"

func canConstruct(s string, k int) bool {
	if len(s) < k {
		return false
	}

	hashMap := make([]int, 26)
	for _, char := range s {
		hashMap[char-'a']++
	}

	odd_char := 0

	for _, value := range hashMap {
		if value%2 != 0 {
			odd_char++
		}
	}

	if odd_char > k {
		return false
	}
	return true
}

func main() {
	fmt.Println(canConstruct("leetcode", 3))
}
