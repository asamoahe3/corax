package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func greetings() string {
	return fmt.Sprintf("Good day 😇, its %v on our clock", time.Now().Format(time.RFC1123Z))
}

// func homeHandler() writes the current time when initial request is made
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte(greetings()))
}

// func dummyProduct() return a json endcoded data from dummyjson.com
func dummyProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
