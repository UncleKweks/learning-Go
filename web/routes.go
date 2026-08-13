package main

import (
	"net/http"
	"path/filepath"
	//"time"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	staticDir := filepath.Join(".", "web", "public")
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir(staticDir))))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/about", app.about)
	mux.HandleFunc("/contact", app.contact)

	return mux
}