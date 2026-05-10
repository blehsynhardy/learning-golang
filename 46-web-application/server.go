package main

import (
	"net/http"
	"time"
)

func (app *Application) serve() error {

	srv := &http.Server{
		Addr:        ":4000",
		ReadTimeout: 2 * time.Second,
		Handler:     app.routes(),
	}

	return srv.ListenAndServe()
}
