package movies

import (
	"bytes"
	"encoding/json"
	"github.com/yodra/awesome-golang-formation/pkg/server/handler"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/yodra/awesome-golang-formation/pkg/storage/storagemocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

func TestCreateHandler(t *testing.T) {
	endpoint := "/movies"
	movieRepository := new(storagemocks.MovieRepository)
	movieRepository.On("Save", mock.Anything, mock.AnythingOfType("domain.Movie")).Return(nil)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST(endpoint, CreateHandler(movieRepository))

	t.Run("should be return StatusCreated", func(t *testing.T) {
		createMovieRequest := handler.CreateMovieRequest{
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

		if status := recorder.Code; status != http.StatusCreated {
			t.Errorf("se ha roto pollito got: %v\n want: %v", status, http.StatusCreated)
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
