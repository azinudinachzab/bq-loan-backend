package middlewares

import (
	"errors"
	"log"
	"net/http"

	"github.com/azinudinachzab/bq-loan-backend/api/auth"
	"github.com/azinudinachzab/bq-loan-backend/api/responses"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestLogger(r)
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestLogger(r)
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}
		next(w, r)
	}
}

func requestLogger(r *http.Request) {
	log.Printf(`url path: %v`, r.URL)
	log.Printf(`request header: %v`, r.Header)
	log.Printf(`request body: %+v`, r.Body)
}
