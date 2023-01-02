package handler

import (
	"fmt"
	"net/http"
	"time"

	"peerac/go-chi-rest-example/app/helper"
	"peerac/go-chi-rest-example/app/logger"
	"peerac/go-chi-rest-example/internal/data"
)

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string   `json:"title"`
		Year    int32    `json:"year"`
		Runtime int32    `json:"runtime"`
		Genres  []string `json:"genres"`
	}

	err := helper.ReadJSON(w, r, &input)
	if err != nil {
		logger.ErrorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func ShowMovie(w http.ResponseWriter, r *http.Request) {
	id, err := helper.ReadIDParam(r)
	if err != nil {
		logger.NotFoundResponse(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Tenet",
		Year:      2022,
		Runtime:   130,
		Genres:    []string{"action", "war", "thriller"},
	}

	err = helper.WriteJSON(w, http.StatusOK, helper.Envelope{"movie": movie}, nil)
	if err != nil {
		logger.ServerErrorResponse(w, r, err)
	}
}

func ListMovies(w http.ResponseWriter, r *http.Request) {

}

func EditMovie(w http.ResponseWriter, r *http.Request) {

}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {

}
