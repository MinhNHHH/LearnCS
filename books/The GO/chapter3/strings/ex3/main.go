package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(isSameLetters("abc", "bac"))
}

func sortString(s string) string {
	bytes := []byte(s)
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})
	return string(bytes)
}

func isSameLetters(s1, s2 string) bool {
	s1 = sortString(s1)
	s2 = sortString(s2)
	return s1 == s2
}
