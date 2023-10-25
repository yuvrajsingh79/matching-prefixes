package app

import (
	"fmt"

	cache "github.com/yuvrajsingh79/matching-prefixes/pkg/controller"
)

func Run() {
	// // Initialize configurations
	// config.Load()

	// // Initialize cache
	cache.Init()

	// // Initialize and start the server
	// router := server.Init()

	port := ":8080"
	fmt.Printf("Server is running on port %s\n", port)
	// if err := http.ListenAndServe(port, router); err != nil {
	// 	log.Fatalf("Server failed to start: %v", err)
	// }
}
