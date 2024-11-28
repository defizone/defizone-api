package middleware

import (
	"log"
	"net/http"
	"time"
)

var API_TOKEN = "123"

func LoggingMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf(r.Method+" "+r.URL.Path, time.Since(start))
	}
}

func AuthMiddleware(next func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Defizone-Token")
		if token == API_TOKEN {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
	}
}
