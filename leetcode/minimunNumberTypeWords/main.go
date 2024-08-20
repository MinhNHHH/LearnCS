package main

import (
	"fmt"
	"sort"
)

func minimumPushes(word string) int {
	set := make([]int, 26)
	res := 0
	press := 1
	taken := 0

	for _, char := range word {
		set[int(char-'a')]++
	}
	sort.Slice(set, func(i, j int) bool { return set[i] > set[j] })

	for key := range set {
		res += press * set[key]
		taken++
		if taken == 8 {
			press++
			taken = 0
		}
	}
	return res
}

func main() {
	fmt.Println(minimumPushes("aabbccddeeffgghhiiiiii"))
	fmt.Println(minimumPushes("afhtgpque"))
}
