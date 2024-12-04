package main

import "fmt"

func incrementString(s byte) rune {
	r := rune(s)
	if r == 'z' {
		r = 'a'
	} else {
		r++
	}
	return r
}

func canMakeSubsequence(str1 string, str2 string) bool {
	i, j := 0, 0

	for i < len(str1) && j < len(str2) {
		if str1[i] == str2[j] {
			j++
		} else {
			newS := incrementString(str1[i])
			if byte(newS) == str2[j] {
				j++
			}
		}
		i++
	}
	if j == len(str2) {
		return true
	}
	return false
}

func main() {
	fmt.Println(canMakeSubsequence("zc", "ad"))
}
