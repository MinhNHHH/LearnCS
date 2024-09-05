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

	// mux.Use(app.enableCORS)

	// authenticate routes - auth handler, refresh
	mux.Post("/auth", app.authenticate)
	mux.Post("/refresh-token", app.refresh)

	// test handler
	mux.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		var payload = struct {
			Message string `json:"message"`
		}{
			Message: "hello world",
		}
		_ = app.writeJSON(w, http.StatusOK, payload)
	})
	// protected routes
	mux.Route("/users", func(muxUser chi.Router) {
		// use auth middleware
		muxUser.Get("/", app.allUsers)
		muxUser.Get("/{userId}", app.getUser)
		muxUser.Delete("/{userId}", app.deleteUser)
		muxUser.Put("/", app.insertUser)
		muxUser.Patch("/", app.updateUser)
	})

	return mux
}
