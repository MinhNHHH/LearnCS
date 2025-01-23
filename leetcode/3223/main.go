package main

import "fmt"

func minimumLength(s string) int {
	mapStr := make([]int, 26)

	for _, char := range s {
		mapStr[char-'a']++
	}

	res := 0
	for _, value := range mapStr {
		for value >= 3 {
			value -= 2
		}
		res += value
	}
	return res
}

func main() {
	fmt.Println(minimumLength("abaacbcbb"))
}
