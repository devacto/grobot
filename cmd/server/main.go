package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	mux := http.NewServeMux()

	// search allows you to search based on keywords
	// http://localhost:5000/search?q=test1+test2
	// q is the query parameter
	mux.HandleFunc("/search", search)

	// starting up the server
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	server.ListenAndServe()
}
