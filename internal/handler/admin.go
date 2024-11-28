package handler

import "net/http"

func CreateLicense(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test"))
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
