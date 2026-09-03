package main

import (
	"net/http"
)

func (app *Application) render(w http.ResponseWriter, r *http.Request, templateName string, data *templateData) {

	if app.rc == nil {

		http.Error(w, "render cache not initialized", http.StatusInternalServerError)

		return

	}

	app.rc.Render(w, templateName, app.defaultTemplateData(data, r))

}

func (app *Application) defaultTemplateData(data *templateData, r *http.Request) *templateData {

	if data == nil {
		data = &templateData{}
	}

	data.Flash = app.session.PopString(r, "flash")
	data.IsAuthenticated = app.isAuthenticated(r)

	return data

}
