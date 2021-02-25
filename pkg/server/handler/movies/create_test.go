package movies

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
		require.NoError(t, err)


		req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(bodyRequest))
		require.NoError(t, err)

		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, req)
		result := recorder.Result()

		assert.Equal(t, http.StatusCreated, result.StatusCode)
	})

	t.Run("should be return StatusBadRequest when the request it is wrong", func(t *testing.T) {
		createMovieRequest := `{
			"Name": "Peliculon",
			"anio": "tueni",
		}`

		req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer([]byte(createMovieRequest)))
		require.NoError(t, err)

		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, req)
		result := recorder.Result()

		assert.Equal(t, http.StatusBadRequest, result.StatusCode)
	})
}
