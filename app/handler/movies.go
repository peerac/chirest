package handler

import (
	"fmt"
	"net/http"

	"peerac/go-chi-rest-example/app/helper"
)

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

func ShowMovie(w http.ResponseWriter, r *http.Request) {
	id, err := helper.ReadIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	// Otherwise, interpolate the movie ID in a placeholder response.
	fmt.Fprintf(w, "show the details of movie %d\n", id)
}

func ListMovies(w http.ResponseWriter, r *http.Request) {

}

func EditMovie(w http.ResponseWriter, r *http.Request) {

}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {

}
