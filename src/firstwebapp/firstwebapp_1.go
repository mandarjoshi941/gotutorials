package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>This is Title</h1>")
	fmt.Fprintf(w, "<p>This is para</p>")
	fmt.Fprintf(w, "<p>This is %s </p>", "END")
}
func main() {
	http.HandleFunc("/", index_handler)
	http.ListenAndServe("localhost:8000", nil)
}
