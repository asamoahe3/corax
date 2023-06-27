package main

import (
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	cor := cors.New(cors.Options{
		AllowedOrigins: []string{"http://127.0.0.1:5501"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"GET"},
	})

	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/dummyjson", dummyProduct)

	log.Fatal(http.ListenAndServe(":"+port, cor.Handler(mux)))
}
