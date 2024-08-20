package main

import (
	"fmt"
)

func lengthOfLastWord(s string) int {
	res := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 32 {
			if res == 0 {
				continue
			} else {
				return res
			}
		}
		res++
	}
	return res
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
}
