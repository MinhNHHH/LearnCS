package main

import (
	"fmt"
)

// You are given a string s consisting only of characters 'a' and 'b'​​​​.
// You can delete any number of characters in s to make s balanced.
// s is balanced if there is no pair of indices (i,j) such that i < j and s[i] = 'b' and s[j]= 'a'.
// Return the minimum number of deletions needed to make s balanced.

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minimumDeletions(s string) int {
	b_count := 0
	res := 0
	for _, char := range s {
		if string(char) == "a" {
			res = min(res, b_count)
		} else {
			b_count++
		}
	}
	return res
}

func main() {
	fmt.Println(minimumDeletions("aababbab"))
}
