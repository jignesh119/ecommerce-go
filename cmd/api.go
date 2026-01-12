package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// mount
// can use chi/fibre/gorilla mux, etc routers/
func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// middleware
	r.Use(middleware.RequestID) // required for rate-limiting
	r.Use(middleware.RealIP)    // also used in rate-limiting, analytics, tracing
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hit /"))
	})

	// http.ListenAndServe(":3333", r)
	return r
}

// run
func (app *application) run(h http.Handler) error {
	// gracefulll shutdown script
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}
	fmt.Printf("Server started at %s", app.config.addr)
	return srv.ListenAndServe()
}

// for dependency injection
type application struct {
	config config
	// logger
	// db driver
}

// configuration of server
type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string // user= passwor= dbname=
}
