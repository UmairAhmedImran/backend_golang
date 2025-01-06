package main

import (
	"log"

	"github.com/UmairAhmedImran/backend_golang/internal/db"
	"github.com/UmairAhmedImran/backend_golang/internal/env"
	"github.com/UmairAhmedImran/backend_golang/internal/store"
)

func main() {

	db_addr := "postgresql://admin:nTsZUv7vT9J7yZcxxqGZJPoH4yHM4npf@dpg-ctto1trv2p9s738k6s1g-a.oregon-postgres.render.com/socialnetwork_plkf"

	cfg := config {
            addr: env.GetString("ADDR", ":8000"),
			db: dbConfig{
				addr:        env.GetString("DB_ADDR", db_addr),
                maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
                maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
                maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
			},
        }

	db, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("database connection pool established")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store: store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
