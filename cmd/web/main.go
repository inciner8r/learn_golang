package main

import (
	"net/http"

	"github.com/inciner8r/simple-go-service/pkg/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.ListenAndServe(":3000", nil)
}
