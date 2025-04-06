package main

import (
	"log"

	"github.com/MdHasib01/go_social_app/internal/env"
	"github.com/MdHasib01/go_social_app/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":5080"),
	}
	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
