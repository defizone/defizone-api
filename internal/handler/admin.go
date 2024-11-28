package handler

import (
	"net/http"
)

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
