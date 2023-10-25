package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux" // Import the gorilla/mux package
	"github.com/yuvrajsingh79/matching-prefixes/pkg/controller"
)

// Run starts the Matching-Prefixes application.
func Run() {
	port := "8080" // Set your desired port here
	fmt.Printf("Server is running on port %s\n", port)

	// Initialize the cache
	controller.Init()

	router := mux.NewRouter()
	router.HandleFunc("/prefix-match/{input}", controller.HandlePrefixMatch)

	if err := http.ListenAndServe(":"+port, router); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
