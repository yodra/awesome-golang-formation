package hello

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.SecureJSON(http.StatusOK, "Hello baby!! 👋")
	}
}
