package controller

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

// Trie represents a prefix tree.
type Trie struct {
	root *TrieNode
}

// NewTrie creates a new Trie.
func NewTrie() *Trie {
	return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

// Insert inserts a prefix into the Trie.
func (t *Trie) Insert(word string) {
	node := t.root
	for _, char := range word {
		if _, exists := node.children[char]; !exists {
			node.children[char] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[char]
	}
	node.isEnd = true
}

// FindLongestPrefix finds the longest matching prefix in the Trie for the input.
func (t *Trie) FindLongestPrefix(input string) string {
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
