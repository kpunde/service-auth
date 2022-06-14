package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"serviceAuth/backend/log"
	"serviceAuth/backend/service"
	"serviceAuth/backend/utility"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		reqContext := utility.RequestContextHandler(context)
		const BEARER_SCHEMA = "Bearer "
		authHeader := context.GetHeader("Authorization")
		if authHeader == "" {
			errorWithAbort(context, "Header not found")
			return
		}
		tokenString := authHeader[len(BEARER_SCHEMA):]

		token, err := service.NewJWTService().ValidateToken(tokenString)
		if token.Valid {
			repo := service.NewPersonService()
			claims := token.Claims.(jwt.MapClaims)
			id := fmt.Sprintf("%v", claims["uid"])
			result, err := repo.FindById(reqContext, id)
			if err != nil {
				errorWithAbort(context, err)
			}

			context.Set("current-user", result.Id)

		} else {
			errorWithAbort(context, err)
		}
	}
}

func errorWithAbort(context *gin.Context, err interface{}) {
	log.Error(err)
	context.AbortWithStatus(http.StatusUnauthorized)
}
