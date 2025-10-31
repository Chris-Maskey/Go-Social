package main

import "log"

func main() {
	cfg := Config{
		Addr: ":8080",
	}

	app := &Application{
		Config: cfg,
	}

	mux := app.Mount()

	log.Fatal(app.Run(mux))
}
