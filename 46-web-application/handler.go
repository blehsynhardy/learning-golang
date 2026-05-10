package main

import (
	"net/http"
)



func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, "index.html", nil)
}

func (app *Application) about(w http.ResponseWriter, r *http.Request) {
	app.render(w, "about.html", nil)
}


