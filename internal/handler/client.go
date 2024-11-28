package handler

import "net/http"

func InitializationApp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
}

func Authentication(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
}

func CheckToken(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
}
