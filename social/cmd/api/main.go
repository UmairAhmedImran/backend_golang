package main

import (
    "log"
	"github.com/UmairAhmedImran/backend_golang/internal/env"
	"github.com/UmairAhmedImran/backend_golang/internal/store"
)

func main() {

	cfg := config {
            addr: env.GetString("ADDR", ":8000"),
        }

	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store: store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
