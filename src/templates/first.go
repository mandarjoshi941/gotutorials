package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type car struct {
	Name string
	Make string
}

func countriesHandler(w http.ResponseWriter, r *http.Request) {
	car := car{Name: "Swift", Make: "Maruti"}
	t, _ := template.ParseFiles("first.html")
	t.Execute(w, car)
}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hey this is index</h1>")
	fmt.Fprintf(w, "<a href='/countries'>Click here</a>")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/countries", countriesHandler)
	http.ListenAndServe("localhost:8000", nil)
}
