package movies

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/yodra/awesome-golang-formation/server"
	"net/http"
	"net/http/httptest"
	"testing"
)

type CreateMovieRequest struct {
	name   string
	year   string
	author string
}

func TestCreateHandler(t *testing.T) {
	endpoint := "/movies"
	createHandler := CreateHandler(&MockRepository{})

	r := mux.NewRouter()
	r.HandleFunc(endpoint, createHandler).Methods(http.MethodPost)

	createMovieRequest := CreateMovieRequest{
		name:   "Peliculon",
		year:   "tueni tueni",
		author: "yo mismo",
	}

	bodyRequest, err := json.Marshal(createMovieRequest)
	if err != nil {
		t.Fatalf("error marshal request %v", err)
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
}

type MockRepository struct{}

func (repo *MockRepository) Save(_ server.Movie) error {
	return nil
}
