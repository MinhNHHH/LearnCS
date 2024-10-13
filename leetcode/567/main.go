package main

import (
	"fmt"
)

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	hash1 := make([]int, 26)
	hash2 := make([]int, 26)

	for i := 0; i < len(s1); i++ {
		hash1[s1[i]-'a']++
		hash2[s2[i]-'a']++
	}
	fmt.Println(hash1 == hash2)
	// for i := 0; i < len(s2); i++ {
	// 	if hash1
	// }
	fmt.Println(hash1)
	return false
}

func main() {
	fmt.Println(checkInclusion("ab", "eidboaooo"))
}
