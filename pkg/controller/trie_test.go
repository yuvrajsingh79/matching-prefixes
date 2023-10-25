package controller

import (
	"bufio"
	"os"
	"testing"
)

func TestTrie_Insert(t *testing.T) {
	// Create a new Trie.
	tr := NewTrie()

	// Open the prefixes_test.txt file.
	file, err := os.Open("../../testfile.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	// Read the prefixes from the file and insert them into the Trie.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		prefix := scanner.Text()
		tr.Insert(prefix)
	}

	// Test the existence of inserted prefixes.
	tests := []struct {
		prefix     string
		shouldBeIn bool
	}{
		{"apple", true},
		{"app", true},
		{"banana", true},
		{"bat", true},
		{"ball", true},    // Should be in the Trie.
		{"apricot", true}, // Should be in the Trie.
		{"grape", false},
		{"cherry", false},
	}

	for _, test := range tests {
		t.Run(test.prefix, func(t *testing.T) {
			inTrie := tr.FindLongestPrefix(test.prefix) != ""
			if inTrie != test.shouldBeIn {
				t.Errorf("Expected %v for prefix %s, but got %v", test.shouldBeIn, test.prefix, inTrie)
			}
		})
	}
}

func TestTrie_FindLongestPrefix(t *testing.T) {
	// Create a new Trie.
	tr := NewTrie()

	// Open the prefixes_test.txt file.
	file, err := os.Open("../../testfile.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	// Read the prefixes from the file and insert them into the Trie.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		prefix := scanner.Text()
		tr.Insert(prefix)
	}

	// Test the FindLongestPrefix function.
	tests := []struct {
		input    string
		expected string
	}{
		{"apples", "apple"},
		{"banana", "banana"},
		{"batman", "bat"},
		{"apricot", "apricot"},
		{"grapes", ""},
		{"cherry", ""},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			longestPrefix := tr.FindLongestPrefix(test.input)
			if longestPrefix != test.expected {
				t.Errorf("Expected %s for input %s, but got %s", test.expected, test.input, longestPrefix)
			}
		})
	}
}
