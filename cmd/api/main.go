package main

import (
	"log"

	"github.com/Chris-Maskey/go-socials/internal/env"
	"github.com/Chris-Maskey/go-socials/internal/store"
)

func main() {
	cfg := Config{
		Addr: env.GetString("ADDR", ":4200"),
	}

	store := store.NewStorage(nil)

	app := &Application{
		Config: cfg,
		Store:  store,
	}

	mux := app.Mount()

	log.Fatal(app.Run(mux))
}
