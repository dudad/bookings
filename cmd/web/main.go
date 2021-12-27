package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/dudad/bookings/pkg/config"
	"github.com/dudad/bookings/pkg/handlers"
	"github.com/dudad/bookings/pkg/render"
)

var app config.AppConfig
var session *scs.SessionManager

func main() {
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tempCache, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal(err)
	}
	app.TemplateCache = tempCache

	r := handlers.NewRepo(&app)
	handlers.NewHandlers(r)
	render.SetConfig(&app)

	srv := &http.Server{
		Addr:    ":8081",
		Handler: routes(&app),
	}

	srv.ListenAndServe()
}
