package chapter4

import (
	"sort"
)

// reverse reverses a slice of ints in place.
func Reverses(s *[5]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func Rotate(array []int, x int) []int {
	tmp := array[:x]
	array = array[x:]
	array = append(array, tmp[:]...)
	return array
}

func RemoveStringDuplicate(s []string) []string {
	sort.Strings(s)
	i := 0
	for i < len(s)-1 {
		if s[i] == s[i+1] {
			s = append(s[:i], s[i+1:]...)
		} else {
			i += 1
		}
	}
	return s
}
