package handler

import (
	"errors"
	"fmt"
	"net/http"
	"peerac/go-chi-rest-example/internal/validator"

	"peerac/go-chi-rest-example/app/helper"
	"peerac/go-chi-rest-example/app/logger"
	"peerac/go-chi-rest-example/internal/data"

	"gorm.io/gorm"
)

func (h *Handler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string `json:"title"`
		Year    int32  `json:"year"`
		Runtime int32  `json:"runtime"`
	}

	err := helper.ReadJSON(w, r, &input)
	if err != nil {
		logger.BadRequestResponse(w, r, err)
		return
	}

	movie := &data.Movie{
		Title:   input.Title,
		Year:    input.Year,
		Runtime: input.Runtime,
	}

	err = h.models.Movies.AddMovie(movie)
	if err != nil {
		logger.ServerErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/movies/%d", movie.ID))

	err = helper.WriteJSON(w, http.StatusCreated, helper.Envelope{"movie": movie}, headers)
	if err != nil {
		logger.ServerErrorResponse(w, r, err)
		return
	}
}

func (h *Handler) ShowMovie(w http.ResponseWriter, r *http.Request) {
	id, err := helper.ReadIDParam(r)
	if err != nil {
		logger.NotFoundResponse(w, r)
		return
	}

	movie, err := h.models.Movies.FetchMovieByID(id)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			logger.NotFoundResponse(w, r)
		default:
			logger.ServerErrorResponse(w, r, err)
		}
		return
	}

	err = helper.WriteJSON(w, http.StatusOK, helper.Envelope{"movie": movie}, nil)
	if err != nil {
		logger.ServerErrorResponse(w, r, err)
		return
	}
}

func (h *Handler) ListMovies(w http.ResponseWriter, r *http.Request) {
	var filters data.Filters
	v := validator.New()

	qs := r.URL.Query()
	filters.Page = helper.ReadInt(qs, "page", 1)
	filters.PageSize = helper.ReadInt(qs, "page_size", 20)
	filters.Sort = helper.ReadString(qs, "sort", "id")
	filters.Order = helper.ReadString(qs, "order", "asc")
	filters.SortSafeList = []string{"id", "title", "year", "runtime"}
	filters.OrderSafeList = []string{"asc", "desc"}

	if data.ValidateFilters(v, filters); !v.Valid() {
		logger.FailedValidationResponse(w, r, v.Errors)
	}

	movies, err := h.models.Movies.FetchMovies(filters)
	if err != nil {
		logger.ServerErrorResponse(w, r, err)
		return
	}

	err = helper.WriteJSON(w, http.StatusOK, helper.Envelope{"movies": movies}, nil)
	if err != nil {
		logger.ServerErrorResponse(w, r, err)
		return
	}
}

func (h *Handler) EditMovie(w http.ResponseWriter, r *http.Request) {
	id, err := helper.ReadIDParam(r)
	if err != nil {
		logger.NotFoundResponse(w, r)
		return
	}

	movie, err := h.models.Movies.FetchMovieByID(id)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			logger.NotFoundResponse(w, r)
		default:
			logger.ServerErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		Title   string `json:"title"`
		Year    int32  `json:"year"`
		Runtime int32  `json:"runtime"`
	}

	err = helper.ReadJSON(w, r, &input)
	if err != nil {
		logger.BadRequestResponse(w, r, err)
		return
	}

	movie.Title = input.Title
	movie.Year = input.Year
	movie.Runtime = input.Runtime

	err = h.models.Movies.UpdateMovie(movie)
	if err != nil {
		logger.ServerErrorResponse(w, r, err)
		return
	}

	err = helper.WriteJSON(w, http.StatusOK, helper.Envelope{"message": "movie successfully updated"}, nil)
	if err != nil {
		logger.ServerErrorResponse(w, r, err)
		return
	}
}

func (h *Handler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	id, err := helper.ReadIDParam(r)
	if err != nil {
		logger.NotFoundResponse(w, r)
		return
	}

	err = h.models.Movies.DeleteMovie(id)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			logger.NotFoundResponse(w, r)
		default:
			logger.ServerErrorResponse(w, r, err)
		}
		return
	}

	err = helper.WriteJSON(w, http.StatusOK, helper.Envelope{"message": "movie successfully deleted"}, nil)
	if err != nil {
		logger.ServerErrorResponse(w, r, err)
		return
	}

}
