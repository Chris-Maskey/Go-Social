package main

import (
	"log"

	"github.com/Chris-Maskey/go-socials/internal/env"
)

func main() {
	cfg := Config{
		Addr: env.GetString("ADDR", ":4200"),
	}

	app := &Application{
		Config: cfg,
	}

	mux := app.Mount()

	log.Fatal(app.Run(mux))
}
