package main

import (
	"log"
	"net/http"
)

func main() {
	// http.FileServer is a handler that serves HTTP requests
	// with the contents of the file system.
	// We tell it to serve files from the current directory (".").
	fs := http.FileServer(http.Dir("."))

	// Handle all requests using the file server.
	http.Handle("/", fs)

	// Log a message that the server is starting.
	log.Println("Starting server at http://localhost:8080")

	// Start the server on port 8080.
	log.Fatal(http.ListenAndServe(":8080", nil))
}
