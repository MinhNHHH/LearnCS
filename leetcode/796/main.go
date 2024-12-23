package main

import "fmt"

func rotate(s string, k int) string {
	if len(s) == 0 {
		return s
	}
	k = k % len(s)
	return s[k:] + s[:k]
}

func rotateString(s string, goal string) bool {
	i := 0

	for i < len(s) {
		s = rotate(s, 1)
		if s == goal {
			return true
		}
		i++
	}

	return false
}

func main() {
	fmt.Println(rotateString("abcde", "abced"))
}
