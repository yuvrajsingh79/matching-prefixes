package controller

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
)

var prefixTrie *Trie
var cache *Cache

func init() {
	// Initialize the prefixTrie and read prefixes from the file
	prefixTrie = NewTrie()

	// Get the current working directory.
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Create a new file path using the current working directory and the file name "prefixes.txt".
	filePath := filepath.Join(cwd, "../testfile.txt")

	// Check if the file exists.
	if _, err := os.Stat(filePath); err != nil {
		// File does not exist.
		fmt.Println("File does not exist.")
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening prefixes file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		prefix := strings.TrimSpace(scanner.Text())
		prefixTrie.Insert(prefix)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading prefixes file:", err)
	}
}

// HandlePrefixMatch is an HTTP handler that matches the longest prefix for a given input.
func HandlePrefixMatch(w http.ResponseWriter, r *http.Request, results chan string) {
	vars := mux.Vars(r)
	input := vars["input"]

	// Calculate the matching prefix
	matchingPrefix := prefixTrie.FindLongestPrefix(input)

	// Send the matching prefix to the results channel
	results <- matchingPrefix

	if matchingPrefix == "" {
		http.Error(w, "No matching prefix found", http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(matchingPrefix))
	}
}

// StartServer starts the HTTP server on the specified port.
func StartServer(port string, c *Cache) error {
	router := mux.NewRouter()
	results := make(chan string)

	cache = c // Initialize the cache

	router.HandleFunc("/prefix-match/{input}", func(w http.ResponseWriter, r *http.Request) {
		go HandlePrefixMatch(w, r, results)
	})

	addr := fmt.Sprintf(":%s", port)

	go func() {
		// This goroutine listens for results on the channel and caches the matching prefixes.
		for result := range results {
			if result != "No matching prefix found" {
				// Cache the matching prefix
				cache.Set(result, result)
			}
		}
	}()

	return http.ListenAndServe(addr, router)
}
