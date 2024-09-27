package main

import "fmt"

type TrieNode struct {
	children map[rune]*TrieNode
	count    int
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: map[rune]*TrieNode{},
		},
	}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, char := range word {
		if _, exists := node.children[char]; !exists {
			node.children[char] = &TrieNode{children: map[rune]*TrieNode{}}
		}
		node = node.children[char]
		node.count++
	}
}

func (t *Trie) GetPrefixScore(word string) int {
	score := 0
	node := t.root
	for _, char := range word {
		node = node.children[char]
		score += node.count
	}
	return score
}

func sumPrefixScores(words []string) []int {
	res := make([]int, len(words))
	trie := NewTrie()

	for _, word := range words {
		trie.Insert(word)
	}

	for index, word := range words {
		res[index] = trie.GetPrefixScore(word)
	}
	return res
}

func main() {
	fmt.Println(sumPrefixScores([]string{"abc", "ab", "bc", "b"}))
}
