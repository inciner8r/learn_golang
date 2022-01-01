package main

import (
	"fmt"
	"net/http"
)

func rootHandler(rw http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(rw, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(rw, "unsupported method", http.StatusNotFound)
		return
	}

	fmt.Fprintf(rw, "kekw")
}

func formHandler(rw http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprint(rw, "parseForm er: %v", err)
		return
	}
	fmt.Printf("post successful")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(rw, "name:%s \naddress:%s", name, address)
}

func main() {
	FileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", FileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", rootHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
