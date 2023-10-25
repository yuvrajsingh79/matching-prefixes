package controller

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux" // Import the gorilla/mux package
)

var prefixTrie *Trie

// Configuration for the prefixes file path.
const prefixesFilePath = "../../prefixes.txt"

func init() {
	// Initialize the prefixTrie and read prefixes from the file
	prefixTrie = NewTrie()

	file, err := os.Open(prefixesFilePath)
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

	matchingPrefix := prefixTrie.FindLongestPrefix(input)
	if matchingPrefix == "" {
		http.Error(w, "No matching prefix found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(matchingPrefix))
}

// StartServer starts the HTTP server on the specified port.
func StartServer(port string) error {
	router := mux.NewRouter()
	router.HandleFunc("/prefix-match/{input}", HandlePrefixMatch)

	addr := fmt.Sprintf(":%s", port)
	return http.ListenAndServe(addr, router)
}
