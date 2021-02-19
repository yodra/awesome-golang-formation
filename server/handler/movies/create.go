package movies

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/yodra/awesome-golang-formation/server"
)

type HttpHandler func(w http.ResponseWriter, _ *http.Request)

func CreateHandler(repository server.MovieRepo) HttpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("could not read body from request: %v", err)
			http.Error(w, server.WriteErrorJSON(http.StatusBadRequest, err.Error()), http.StatusBadRequest)
			return
		}

		var movieRequest server.CreateMovieRequest
		err = json.Unmarshal(reqBody, &movieRequest)
		if err != nil {
			log.Printf("could not parse from body data: %v", err)
			http.Error(w, server.WriteErrorJSON(http.StatusBadRequest, err.Error()), http.StatusBadRequest)
			return
		}

		movie, err:= server.NewMovie(movieRequest.Name, movieRequest.Year, movieRequest.Author)
		if err != nil {
			log.Printf("could not create the movie domain: %v", err)
			http.Error(w, server.WriteErrorJSON(http.StatusExpectationFailed, err.Error()), http.StatusExpectationFailed)
			return
		}

		err = repository.Save(ctx, movie)
		if err != nil {
			log.Printf("could not save the movie: %v", err)
			http.Error(w, server.WriteErrorJSON(http.StatusExpectationFailed, err.Error()), http.StatusExpectationFailed)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
