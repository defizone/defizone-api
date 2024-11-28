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

	// Client routes
	router.HandleFunc("/client/v1/initialization_app", handler.InitializationApp)
	router.HandleFunc("/client/v1/authentication", handler.Authentication)
	router.HandleFunc("/client/v1/check_token", handler.CheckToken)

	// Admin routes
	router.HandleFunc("GET /_/admin/license/create_license", middleware.AuthMiddleware(handler.CreateLicense))
	router.HandleFunc("/_/admin/license/update_license", handler.UpdateLicense)
	router.HandleFunc("/_/admin/license/delete_license", handler.DeleteLicense)
	router.HandleFunc("/_/admin/license/all_license", handler.AllLicense)

	server := &http.Server{
		Addr:    s.addr,
		Handler: middleware.LoggingMiddleware(router),
	}
	log.Printf("server listening on %s", s.addr)

	return server.ListenAndServe()
}
