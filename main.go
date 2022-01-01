package main

import (
	"fmt"
	"net/http"
)

func rootHandler(rw http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(rw, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(rw, "unsupported method", http.StatusNotFound)
		return
	}

	fmt.Fprintf(rw, "kekw")
}

func main() {
	http.HandleFunc("/", rootHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
