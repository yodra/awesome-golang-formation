package movies

import (
	"github.com/yodra/awesome-golang-formation/server"
	"net/http"
)

type HttpHandler func(w http.ResponseWriter, _ *http.Request)

func CreateHandler(repository server.MovieRepo) HttpHandler {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("Hello baby!! 👋"))
		if err != nil {
			http.Error(w, server.WriteErrorJSON(http.StatusInternalServerError, err.Error()), http.StatusInternalServerError)
			return
		}
	}
}
