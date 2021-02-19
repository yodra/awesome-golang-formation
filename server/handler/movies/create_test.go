package movies

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/yodra/awesome-golang-formation/server"
	"github.com/yodra/awesome-golang-formation/server/storage/storagemocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateHandler(t *testing.T) {
	endpoint := "/movies"
	createHandler := CreateHandler(&storagemocks.MockRepository{})

	r := mux.NewRouter()
	r.HandleFunc(endpoint, createHandler).Methods(http.MethodPost)

	t.Run("should be return StatusOk", func(t *testing.T) {
		createMovieRequest := server.CreateMovieRequest{
			Name:   "Peliculon",
			Year:   "tueni tueni",
			Author: "yo mismo",
		}

		bodyRequest, err := json.Marshal(createMovieRequest)
		if err != nil {
			t.Fatalf("error on marshal request %v", err)
		}

		req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(bodyRequest))
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, req)

		if status := recorder.Code; status != http.StatusOK {
			t.Errorf("se ha roto pollito got: %v\n want: %v", status, http.StatusOK)
		}
	})
	t.Run("should be return StatusBadRequest when the request it is wrong", func(t *testing.T) {
		createMovieRequest := `{
			"Name": "Peliculon",
			"anio": "tueni",
		}`

		req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer([]byte(createMovieRequest)))
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, req)

		if status := recorder.Code; status != http.StatusBadRequest {
			t.Errorf("error got: %v\n want: %v", status, http.StatusBadRequest)
		}
	})
}
