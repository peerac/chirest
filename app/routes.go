package main

import (
	"net/http"

	"peerac/go-chi-rest-example/app/handler"
	"peerac/go-chi-rest-example/app/helper"
	"peerac/go-chi-rest-example/app/logger"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() *chi.Mux {
	r := chi.NewRouter()
	r.NotFound(logger.NotFoundResponse)

	// healthcheck route
	r.Get("/v1/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		env := helper.Envelope{
			"status": "available",
			"system_info": map[string]string{
				"environment": app.config.Env,
				"version":     app.config.Version,
			},
		}

		err := helper.WriteJSON(w, http.StatusOK, env, nil)
		if err != nil {
			logger.ServerErrorResponse(w, r, err)
		}
	})

	// movies route
	r.Post("/v1/movies", handler.CreateMovie)
	r.Get("/v1/movies/{id}", handler.ShowMovie)

	return r
}
