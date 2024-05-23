package main

import "fmt"

func main() {
	fmt.Println(string(reverseString([]byte("abcd"))))
}

func reverseString(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
