package main

import (
	"defizone-api/internal/handler"
	"defizone-api/internal/middleware"
	"log"
	"net/http"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Start() error {
	router := http.NewServeMux()

	// User routes
	router.HandleFunc("/client/v1/initialization_app", handler.InitializationApp)
	router.HandleFunc("/client/v1/authentication", handler.Authentication)
	router.HandleFunc("/client/v1/check_token", handler.CheckToken)

	// Admin routes
	router.HandleFunc("/_/admin/create_key", handler.CreateLicense)

	server := &http.Server{
		Addr:    s.addr,
		Handler: middleware.ResponseLogging(router),
	}
	log.Printf("server listening on %s", s.addr)

	return server.ListenAndServe()
}
