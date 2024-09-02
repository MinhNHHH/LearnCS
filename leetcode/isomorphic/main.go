package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	hashTable1 := map[rune]byte{}
	hashTable2 := map[rune]byte{}
	if len(s) != len(t) {
		return false
	}

	for index, char := range s {
		hashTable1[char] = t[index]
	}

	for index, char := range t {
		hashTable2[char] = s[index]
	}

	for i := 0; i < len(s); i++ {
		if hashTable1[rune(s[i])] != t[i] {
			return false
		}
		if hashTable2[rune(t[i])] != s[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isIsomorphic("foo", "egg"))
	fmt.Println(isIsomorphic("foo", "bar"))
}
