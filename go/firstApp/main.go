package main

import (
	"log"
	"net/http"

	"encoding/json"
)

type result struct {
	Text string `json:text`
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result{Text: "testだぜ"})
	})
	log.Println("Go API running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
