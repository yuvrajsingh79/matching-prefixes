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

// Configuration for the prefixes file path.
// const prefixesFilePath = "../../prefixes.txt"

func init() {
	// Initialize the prefixTrie and read prefixes from the file
	prefixTrie = NewTrie()

	// Get the current working directory.
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Create a new file path using the current working directory and the file name "myfile.txt".
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
func HandlePrefixMatch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	input := vars["input"]

	// Check if the matching prefix is in the cache
	cache := GetCache()
	cachedPrefix, ok := cache.Get(input)
	if ok {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(cachedPrefix.(string)))
		return
	}

	results := make(chan string)
	defer close(results)

	go func(results chan<- string) {
		// Calculate the matching prefix
		matchingPrefix := prefixTrie.FindLongestPrefix(input)
		if matchingPrefix == "" {
			results <- "No matching prefix found"
		} else {
			results <- matchingPrefix
		}
	}(results)

	// Retrieve the matching prefix from the goroutine
	matchingPrefix := <-results

	// // Calculate the matching prefix
	// matchingPrefix := prefixTrie.FindLongestPrefix(input)
	// if matchingPrefix == "" {
	// 	http.Error(w, "No matching prefix found", http.StatusNotFound)
	// 	return
	// }

	// Cache the matching prefix
	cache.Set(input, matchingPrefix)

	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte(matchingPrefix))

	if matchingPrefix == "No matching prefix found" || matchingPrefix == "" {
		http.Error(w, "No matching prefix found", http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(matchingPrefix))
	}
}

// StartServer starts the HTTP server on the specified port.
func StartServer(port string) error {
	router := mux.NewRouter()
	router.HandleFunc("/prefix-match/{input}", HandlePrefixMatch)

	addr := fmt.Sprintf(":%s", port)
	return http.ListenAndServe(addr, router)
}
