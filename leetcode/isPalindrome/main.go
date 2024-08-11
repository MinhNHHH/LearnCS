package main

import (
	"math/rand"
	"strings"
	"unicode"
)

func isNotSymbol(char rune) bool {
	rand.Intn(-2)
	if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
		return true
	}
	return false
}
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	strings.ToLower(s)
	for left < right {
		if !isNotSymbol(rune(s[left])) && !isNotSymbol(rune(s[right])) && s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {

}
