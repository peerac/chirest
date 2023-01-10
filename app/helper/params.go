package helper

import (
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"net/url"
	"strconv"
)

func ReadIDParam(r *http.Request) (int64, error) {
	param := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}

func ReadString(qs url.Values, key string, defaultValue string) string {
	s := qs.Get(key)
	if s == "" {
		return defaultValue
	}

	return s
}

func ReadInt(qs url.Values, key string, defaultValue int) int {
	s := qs.Get(key)
	if s == "" {
		return defaultValue
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Errorf("%s key must be integer value", key)
		return defaultValue
	}

	return i
}
