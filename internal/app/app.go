package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yuvrajsingh79/matching-prefixes/pkg/controller"
)

// Run starts the Matching-Prefixes application.
func Run() {
	port := "8080" // Set your desired port here
	fmt.Printf("Server is running on port %s\n", port)

	// Initialize the cache
	controller.Init()

	results := make(chan string)
	var errCh = make(chan error, 1)

	router := mux.NewRouter()
	router.HandleFunc("/prefix-match/{input}", func(w http.ResponseWriter, r *http.Request) {
		go controller.HandlePrefixMatch(w, r, results)
	})

	go func() {
		errCh <- http.ListenAndServe(":"+port, router)
	}()

	// Concurrently handle results and errors
	go func() {
		for {
			select {
			case result := <-results:
				// Process results as needed
				fmt.Printf("Matching Result: %s\n", result)

			case err := <-errCh:
				if err != nil {
					fmt.Printf("Server failed to start: %v\n", err)
				}
				return
			}
		}
	}()

	select {
	case <-errCh:
		close(results)
	}
}
