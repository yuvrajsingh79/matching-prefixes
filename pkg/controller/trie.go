package controller

import "sync"

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
	mu   sync.RWMutex // Mutex for thread-safe access
}

func NewTrie() *Trie {
	return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

func (t *Trie) Insert(word string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	node := t.root
	for _, char := range word {
		if _, exists := node.children[char]; !exists {
			node.children[char] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[char]
	}
	node.isEnd = true
}

func (t *Trie) FindLongestPrefix(input string) string {
	t.mu.RLock()
	defer t.mu.RUnlock()

	node := t.root
	longestPrefix := ""
	for _, char := range input {
		if _, exists := node.children[char]; exists {
			longestPrefix += string(char)
			node = node.children[char]
		} else {
			break
		}
	}
	return longestPrefix
}
