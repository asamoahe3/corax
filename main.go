package main

import (
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

var TempJson map[string]interface{}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	cor := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"GET"},
	})

	if err := MakeGetRequest("https://dummyjson.com/users", &TempJson); err != nil {
		log.Fatal(err.Error())
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/dummyjson", dummyProduct)

	log.Fatal(http.ListenAndServe(":"+port, cor.Handler(mux)))
}
