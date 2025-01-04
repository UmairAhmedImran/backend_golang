package main

import (
    "log"
	"github.com/UmairAhmedImran/backend_golang/internal/env"
	"github.com/UmairAhmedImran/backend_golang/internal/store"
)

func main() {

	cfg := config {
            addr: env.GetString("ADDR", ":8000"),
			db: dbConfig{
				addr:        env.GetString("DB_ADDR", "postgres://user:adminpassword@localhost:/social?
				sslmode=disable"),
                maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
                maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
                maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15min"),
			}
        }

	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store: store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
