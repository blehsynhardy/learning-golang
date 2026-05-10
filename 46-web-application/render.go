package main

import (
	"net/http"
)

func (app *Application) render(w http.ResponseWriter, templateName string, data interface{}) {

	if app.rc == nil {

		http.Error(w, "render cache not initialized", http.StatusInternalServerError)

		return

	}

	app.rc.Render(w, templateName, data)

}
