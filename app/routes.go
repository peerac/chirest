package main

import (
	"net/http"

	"peerac/go-chi-rest-example/app/handler"
	"peerac/go-chi-rest-example/app/helper"
	"peerac/go-chi-rest-example/app/logger"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func (app *application) routes(db *gorm.DB) *chi.Mux {
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

	handlers := handler.NewHandler(db)
	// movies route
	r.Post("/v1/movies", handlers.CreateMovie)
	r.Get("/v1/movies", handlers.ListMovies)
	r.Get("/v1/movies/{id}", handlers.ShowMovie)
	r.Patch("/v1/movies/{id}", handlers.EditMovie)
	r.Delete("/v1/movies/{id}", handlers.DeleteMovie)

	return r
}
