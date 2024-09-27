package main

import (
	"fmt"
	"strings"
)

func uncommonFromSentences(s1 string, s2 string) []string {
	res := []string{}
	s1 = s1 + " " + s2
	newS := strings.Split(s1, " ")
	hash := map[string]int{}

	for _, str := range newS {
		hash[str]++
	}

	for str, value := range hash {
		if value == 1 {
			res = append(res, str)
		}
	}
	return res
}

func main() {
	fmt.Println(uncommonFromSentences("abcd def abcd xyz", "ijk def ijk"))
}
