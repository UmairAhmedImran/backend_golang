package main

import (
	"net/http"
	"log"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (app *application) mount() *http.ServeMux{

    mux := http.NewServeMux()

    mux.HandleFunc("/api/users", api.getUsersHandler)
	
    mux.HandleFunc("/api/users", api.createUsersHandler)

    return mux
}

func (app *application) run(mux *http.ServeMux) error {

	srv := &http.Server{
		Addr:    app.config.addr,
		Handler: mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout: time.Second * 10,
		IdleTimeout: time.Minute,
	}

	log.Printf("Server started at %s", app.config.addr)

	return srv.ListenAndServe()
}