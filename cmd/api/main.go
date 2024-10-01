package main

import (
	"log"

	"github.com/careofyou/rest-backend.git/internal/db"
	"github.com/careofyou/rest-backend.git/internal/env"
	"github.com/careofyou/rest-backend.git/internal/store"
)

func main() {
    cfg :=  config{
        addr: env.GetString("ADDR", ":8081"),
        db: dbConfig{
            addr: env.GetString("DB_ADDR", "postgres://user:password@localhost/social?sslmode=disable"),
            maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
            maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
            maxIdleTime: env.GetString("DB_MAX_IDLE_TIME", "15m"),
        },
    }

    db, err := db.New(
        cfg.db.addr,
        cfg.db.maxOpenConns,
        cfg.db.maxIdleConns,
        cfg.db.maxIdleTime,
        )
    if err != nil {
        log.Panic(err)
    }

    defer db.Close()
    log.Println("db connection pull established")

    store := store.NewStorage(db) 
    
    app := &application{
        config: cfg,
        store: store,
    }
    mux := app.mount()
    log.Fatal(app.run(mux))
}