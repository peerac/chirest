package main

import (
	"net/http"

	"peerac/go-chi-rest-example/app/handler"
	"peerac/go-chi-rest-example/app/helper"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() *chi.Mux {
	r := chi.NewRouter()

	// healthcheck route
	r.Get("/v1/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"status":      "available",
			"environment": app.config.Env,
			"version":     app.config.Version,
		}

		err := helper.JSONSerializer(w, http.StatusOK, data, nil)
		if err != nil {
			app.logger.Println(err)
			http.Error(w, "The server encountered a problem", http.StatusInternalServerError)
		}
	})

	// movies route
	r.Post("/v1/movies", handler.CreateMovie)
	r.Get("/v1/movies/{id}", handler.ShowMovie)

	return r
}
