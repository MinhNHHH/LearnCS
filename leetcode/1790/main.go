package main

import "fmt"

func areAlmostEqual(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	if s1 == s2 {
		return true
	}
	indexs := []int{}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			indexs = append(indexs, i)
		}
		if len(indexs) > 2 {
			return false
		}
	}

	return len(indexs) == 0 || (len(indexs) == 2 && s1[indexs[0]] == s2[indexs[1]] && s1[indexs[1]] == s2[indexs[0]])
}

func main() {
	fmt.Println(areAlmostEqual("bank", "kanb"))
}
