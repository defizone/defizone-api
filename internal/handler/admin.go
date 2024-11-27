package handler

import "net/http"

func CreateLicense(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test"))
}
