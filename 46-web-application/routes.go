package main

import (
	"net/http"
)

func (app *Application) routes() http.Handler {
	mux := http.NewServeMux()

	// Rendering static assets

	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir(app.publicDir))))
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/about", app.about)
	return mux

}
