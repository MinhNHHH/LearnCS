package main

import (
	"fmt"
)

type IntSet struct {
	words []uint64
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) Len() int {
	return len(s.words)
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	s.words[word] |= 1 << bit
	fmt.Println(word, bit, s.words[word])
	i := 0
	for i < s.Len() {
		if s.words[i] == uint64(x) {
			s.words = append(s.words[:i], s.words[i+1:]...)
		}
		i++
	}
}

func main() {
	s := &IntSet{}

	// Adding values to the set
	values := []int{1, 65, 128, 300}

	for _, v := range values {
		s.Add(v)
	}
	s.Remove(65)
	// Printing the internal state of IntSet
	fmt.Println("Internal state of IntSet:")
	for i, word := range s.words {
		fmt.Printf("Word %d: %064b\n", i, word)
	}
}
