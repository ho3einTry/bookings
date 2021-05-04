package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ho3einTry/bookings/pkg/config"
	"github.com/ho3einTry/bookings/pkg/handlers"
	"net/http"
)

func routes(config *config.AppConfig) http.Handler {
	//mux := pat.New()
	//mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	//mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	//return mux
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	//mux.Use(WriteToConsole)
	mux.Use(NoSruve)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*",http.StripPrefix("/static/",fileServer))
	return mux
}
