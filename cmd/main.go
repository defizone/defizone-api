package main

import (
	"log"
	"net/http"
)

func getIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", getIndex)

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
