package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	// register middleware
	mux.Use(middleware.Recoverer)
	mux.Use(app.addIPToContext)
	mux.Use(app.Session.LoadAndSave)

	// register routes
	mux.Get("/", app.Home)
	mux.Post("/login", app.Login)
	// static assets
	mux.Route("/user", func(muxUser chi.Router) {
		muxUser.Use(app.auth)
		muxUser.Get("/profile", app.Profile)
		muxUser.Post("/upload-profile-pic", app.UploadProfilePic)
	})
	fileServer := http.FileServer(http.Dir("./../static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}