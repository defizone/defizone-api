package middleware

import (
	"log"
	"net/http"
	"time"
)

var ApiToken = "Ur6PXzDVXWZGgg9dSYRqOQenfJ8EYp"

type ErrorToken struct {
	Status bool `json:"status"`
	Data   struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	} `json:"data"`
}

func LoggingMiddleware(next *http.ServeMux) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf(r.Method+" "+r.URL.Path, time.Since(start))
	}
}

func AuthMiddleware(next func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Defizone-Token")
		if token != ApiToken {
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}
