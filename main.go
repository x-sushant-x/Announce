package main

import (
	"log"
	"net/http"
)

func main() {
	registry := NewRegistry()

	mux := http.NewServeMux()

	mux.HandleFunc("/api/service", makeHandler(registry))

	log.Fatal(http.ListenAndServe(":8080", mux))
}
