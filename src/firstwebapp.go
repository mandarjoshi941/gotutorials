package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get set go!")
}

func info_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contact: +91-8149781104")
}
func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/info", info_handler)
	http.ListenAndServe(":8000", nil)
}
