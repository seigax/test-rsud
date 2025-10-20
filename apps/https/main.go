package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, this is a secure API!")
}

func main() {
	http.HandleFunc("/", helloHandler)

	// Start the HTTPS server
	fmt.Println("Starting HTTPS server on :443...")
	err := http.ListenAndServeTLS("127.0.0.1:443", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
