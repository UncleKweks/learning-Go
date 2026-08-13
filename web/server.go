package main

import (
	"net/http"
	"time"
)

func (app *application) Serve() error {

	srv := &http.Server{
		Addr:     ":8080",
		ReadTimeout: 5 * time.Minute,
		WriteTimeout: 10 * time.Minute,
		ErrorLog: app.errorLog,
		Handler:  app.routes(),
	}
	
	
	return srv.ListenAndServe()
	
}

