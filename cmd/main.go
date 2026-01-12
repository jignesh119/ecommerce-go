package main

import (
	"log"
	"os"
)

// main is entrypoint of whole server
func main() {
	cfg := config{
		addr: ":8080",
		db:   dbConfig{},
	}

	api := application{
		config: cfg,
	}

	if err := api.run(api.mount()); err != nil {
		log.Printf("Server has failed to start, err: %s", err)
		os.Exit(1)
	}
}
