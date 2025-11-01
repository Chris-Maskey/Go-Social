package main

import (
	"log"

	"github.com/Chris-Maskey/go-socials/internal/db"
	"github.com/Chris-Maskey/go-socials/internal/env"
	"github.com/Chris-Maskey/go-socials/internal/store"
)

func main() {
	cfg := Config{
		Addr: env.GetString("ADDR", ":4200"),
		DB: DBConfig{
			Addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable"),
			MaxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			MaxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			MaxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(cfg.DB.Addr, cfg.DB.MaxOpenConns, cfg.DB.MaxIdleConns, cfg.DB.MaxIdleTime)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	log.Println("Database connection pool established")

	store := store.NewStorage(db)

	app := &Application{
		Config: cfg,
		Store:  store,
	}

	mux := app.Mount()

	log.Fatal(app.Run(mux))
}
