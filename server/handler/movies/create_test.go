package movies

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
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
	createHandler := CreateHandler(&MockRepo{})

	r := mux.NewRouter()
	r.HandleFunc(endpoint, createHandler).Methods(http.MethodPost)

	createMovieRequest := CreateMovieRequest{
		name:   "Peliculon",
		year:   "tueni tueni",
		author: "yo mismo",
	}

	bodyRequest, err := json.Marshal(createMovieRequest)
	require.NoError(t, err)
	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(bodyRequest))
	require.NoError(t, err)

	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("se ha roto pollito got: %v want: %v", status, http.StatusOK)
	}
}

type MockRepo struct{}

func (repo *MockRepo) Save(_ server.Movie) error {
	return nil
}
