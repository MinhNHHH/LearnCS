package main

import (
	"fmt"
	"sort"
	"strings"
)

// Node represents a single node in the Trie
type Node struct {
	children map[string]*Node
	isEnd    bool
}

// Trie represents the Trie structure
type Trie struct {
	root *Node
}

// NewTrie initializes a new Trie
func NewTrie() *Trie {
	return &Trie{
		root: &Node{
			children: make(map[string]*Node),
			isEnd:    false,
		},
	}
}

func removeSubfolders(folder []string) []string {
	sort.Slice(folder, func(i, j int) bool {
		return folder[i] < folder[j]
	})
	res := []string{folder[0]}
	j := len(res) - 1
	for i := 1; i < len(folder); i++ {
		if strings.HasPrefix(folder[i], res[j]+"/") {
			continue
		} else {
			res = append(res, folder[i])
			j++
		}
	}

	return res
}

func main() {
	fmt.Println(removeSubfolders([]string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}))
}
