package hello

import (
	"awesome-golang-formation/server"
	"net/http"
)

func Handler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Hello baby!! 👋"))
	if err != nil {
		http.Error(w, server.WriteErrorJSON(http.StatusInternalServerError, err.Error()), http.StatusInternalServerError)
		return
	}
}
