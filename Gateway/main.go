package main

import (
	"encoding/json"
	"net/http"
)

func getPing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("pong")
}

func main() {
	http.HandleFunc("/ping", getPing)
	http.ListenAndServe(":8080", nil)
}
