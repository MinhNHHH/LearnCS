package main

import (
	"fmt"
)

func compressedString(word string) string {
	res := ""
	i := 1
	count := 1
	for i <= len(word) {
		if i == len(word) || word[i] != word[i-1] || count == 9 {
			res = res + fmt.Sprintf("%d%c", count, word[i-1])
			count = 1
		} else {
			count++
		}
		i++
	}
	return res
}

func main() {
	fmt.Println(compressedString("jd"))
}
