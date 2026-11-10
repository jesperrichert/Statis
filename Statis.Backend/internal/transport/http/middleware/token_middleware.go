package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func DefaultMiddleware(ctx *gin.Context) {

	if ctx.GetHeader("Authorization") == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := ctx.GetHeader("Authorization")
	authToken, ok := os.LookupEnv("AUTH_TOKEN")
	if !ok {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	if token != authToken {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Next()
}
