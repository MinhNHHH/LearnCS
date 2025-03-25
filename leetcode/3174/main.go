package main

import (
	"fmt"
	"unicode"
)

func clearDigits(s string) string {

	stacks := []rune{}

	for _, char := range s {
		if !unicode.IsDigit(char) {
			stacks = append(stacks, char)
		} else if len(stacks) > 0 && unicode.IsDigit(char) {
			stacks = stacks[:len(stacks)-1]
		}
	}
	return string(stacks)
}

func main() {
	fmt.Println(clearDigits("cb34"))
}
