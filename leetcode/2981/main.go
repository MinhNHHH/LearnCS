package main

import "fmt"

func main() {
	fmt.Println(maximumLength("aaaa"))
}

func isSpecial(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i-1] != s[i] {
			return false
		}
	}
	return true
}

func maximumLength(s string) int {
	n := len(s)
	hashMap := map[string]int{}

	for i := n - 2; i >= 1; i-- {
		for j := 0; j <= n-i; j++ {
			str := string(s[j : i+j])
			if isSpecial(str) {
				hashMap[str]++
				if hashMap[str] >= 3 {
					return i
				}
			}
		}
	}
	return -1
}
