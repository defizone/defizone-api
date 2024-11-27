package main

import (
	"defizone-api/internal/handler"
	"log"
	"net/http"
)

func main() {
	//TODO: Сделать роутинг для пользователей и администратора
	//TODO: Добавить разделение версий API
	router := http.NewServeMux()

	router.HandleFunc("/", handler.GetIndex)
	router.HandleFunc("/_/admin/create_key", handler.CreateLicense)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Starting server on port 8080")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
