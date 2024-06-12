package main

import (
	"fmt"
)

type TrieNode struct {
	children map[rune]*TrieNode
	failure  *TrieNode
	output   []string
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		failure:  nil,
		output:   []string{},
	}
}

type AhoCorasick struct {
	root *TrieNode
}

func NewAhoCorasick() *AhoCorasick {
	return &AhoCorasick{root: NewTrieNode()}
}

func (ac *AhoCorasick) AddPattern(pattern string) {
	node := ac.root
	for _, char := range pattern {
		if _, ok := node.children[char]; !ok {
			node.children[char] = NewTrieNode()
		}
		node = node.children[char]
	}
	node.output = append(node.output, pattern)
}

func (ac *AhoCorasick) BuildFailureLinks() {
	queue := make([]*TrieNode, 0)
	for _, child := range ac.root.children {
		queue = append(queue, child)
		child.failure = ac.root
	}

	for len(queue) > 0 {
		currNode := queue[0]
		queue = queue[1:]

		for char, child := range currNode.children {
			queue = append(queue, child)
			failure := currNode.failure
			for failure != nil && failure.children[char] == nil {
				failure = failure.failure
			}
			if failure == nil {
				child.failure = ac.root
			} else {
				child.failure = failure.children[char]
			}
			child.output = append(child.output, child.failure.output...)
		}
	}
}

func (ac *AhoCorasick) Search(text string) map[string][]int {
	result := make(map[string][]int)
	node := ac.root
	for i, char := range text {
		for node != nil && node.children[char] == nil {
			node = node.failure
		}
		if node == nil {
			node = ac.root
			continue
		}
		node = node.children[char]
		for _, pattern := range node.output {
			result[pattern] = append(result[pattern], i-len(pattern)+1)
		}
	}
	return result
}

func main() {
	ac := NewAhoCorasick()
	patterns := []string{"he", "she", "his", "hers"}
	for _, pattern := range patterns {
		ac.AddPattern(pattern)
	}
	ac.BuildFailureLinks()
	text := "ushers"
	matches := ac.Search(text)
	fmt.Println("Matches found:")
	for pattern, indices := range matches {
		for _, index := range indices {
			fmt.Printf("Pattern: %s, Index: %d\n", pattern, index)
		}
	}
}
