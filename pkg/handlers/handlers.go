package handlers

import (
	"fmt"
	"net/http"

	"github.com/inciner8r/simple-go-service/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.go.tmpl")
}
func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About page")
}
