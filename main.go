package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home page ")
}
func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About page")
}
func AddValues(x, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.ListenAndServe(":3000", nil)
}
