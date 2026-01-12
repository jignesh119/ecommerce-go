package main

import (
	"log"
	"log/slog"
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

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := api.run(api.mount()); err != nil {
		slog.Error("Server has failed to start", err)
		os.Exit(1)
	}
}
