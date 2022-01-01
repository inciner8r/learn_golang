package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) { fmt.Fprintf(rw, "server up") })
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
