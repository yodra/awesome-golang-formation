package movies

import (
	"github.com/yodra/awesome-golang-formation/pkg/domain"
	"github.com/yodra/awesome-golang-formation/pkg/server/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateHandler(repository domain.MovieRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var movieRequest handler.CreateMovieRequest
		err := ctx.BindJSON(&movieRequest)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		movie, err := domain.NewMovie(movieRequest.Name, movieRequest.Year, movieRequest.Author)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
		}

		err = repository.Save(ctx, movie)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}

		ctx.Status(http.StatusCreated)
	}
}
