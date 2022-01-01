package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("ok")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
