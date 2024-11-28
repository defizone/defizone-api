package handler

import (
	"net/http"
)

type CreateResponse struct {
	Status bool `json:"status"`
	Data   struct {
		ID      string `json:"id"`
		License string `json:"license"`
	} `json:"data"`
}

func CreateLicense(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Beta-ldw9eSYRqOQenfJ8EYp"))
	return
}

func UpdateLicense(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test"))
}

func DeleteLicense(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test"))
}

func AllLicense(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test"))
}
