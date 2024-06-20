package exercises

import "sort"

func Equal(i, j int, s sort.Interface) bool {
	return !s.Less(i, j) && !s.Less(j, i)
}

func IsPalindrome(s sort.Interface) bool {
	max := s.Len() - 1

	for i := 0; i < s.Len()/2; i++ {
		if !Equal(i, max-i, s) {
			return false
		}
	}
	return true
}
