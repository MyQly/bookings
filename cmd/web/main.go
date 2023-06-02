/*
Corresponds to Section 3.25
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/myqly/bookings/pkg/config"
	"github.com/myqly/bookings/pkg/handlers"
	"github.com/myqly/bookings/pkg/render"
)

//const PORTNUMBER = ":8080"

// var sessionManager *scs.SessionManager
var app config.AppConfig

func main() {
	app.Secure = false

	SessionManager := scs.New()
	SessionManager.Lifetime = 24 * time.Hour
	SessionManager.Cookie.Persist = true
	SessionManager.Cookie.SameSite = http.SameSiteLaxMode
	SessionManager.Cookie.Secure = app.Secure

	app.SessionManager = SessionManager

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.UseCache = false
	app.Port = ":8080"
	app.TemplateCache = templateCache

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	/*
		http.HandleFunc("/", handlers.Repo.Home)
		http.HandleFunc("/about", handlers.Repo.About)
		http.HandleFunc("/cache", handlers.Cache)
	*/

	fmt.Printf("Starting server on port (%s) \n", app.Port)
	//_ = http.ListenAndServe(PORTNUMBER, nil)

	srv := &http.Server{
		Addr:    app.Port,
		Handler: app.SessionManager.LoadAndSave(routes(&app)),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
