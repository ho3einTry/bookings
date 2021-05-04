package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/ho3einTry/bookings/pkg/config"
	"github.com/ho3einTry/bookings/pkg/handlers"
	"github.com/ho3einTry/bookings/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig
var sessionManager *scs.SessionManager

func main() {

	app.InProduction = false
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = app.InProduction

	app.SessionManger = sessionManager

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)
	/*	http.HandleFunc("/", handlers.Repo.Home)
		http.HandleFunc("/home", handlers.Repo.Home)
		http.HandleFunc("/about", handlers.Repo.About)*/

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	//_ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}


	err = srv.ListenAndServe()

	log.Fatal(err)

}
