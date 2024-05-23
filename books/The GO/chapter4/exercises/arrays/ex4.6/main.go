package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(squashSpace("Hello,     世界!  This    is a    test."))
}

func squashSpace(s string) string {
	runes := []rune(s)
	res := []rune{}
	index := 0
	subRune := []rune{}
	for index < len(runes) {
		if !unicode.IsSpace(runes[index]) {
			subRune = append(subRune, runes[index])
		} else {
			if len(subRune) > 0 {
				subRune = append(subRune, 32)
				res = append(res, subRune...)
			}
			subRune = []rune{}
		}
		index++
	}
	res = append(res, subRune...)
	return string(res)
}
