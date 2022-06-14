package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"serviceAuth/backend/log"
	"serviceAuth/backend/service"
	"serviceAuth/backend/utility"
)

var (
	_loginService = service.NewLoginService()
	_jwtService   = service.NewJWTService()
)

type credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type tokenResponse struct {
	Token string `json:"token"`
}

func Login(ctx *gin.Context) {
	reqContext := utility.RequestContextHandler(ctx)
	var cred credentials
	err := ctx.ShouldBindJSON(&cred)
	log.Debug(cred)

	if !utility.CheckErrorResponse(ctx, http.StatusBadRequest, err) {
		user, isAuthenticated := _loginService.Login(reqContext, cred.Email, cred.Password)
		if isAuthenticated {
			token := _jwtService.GenerateToken(user.Id)
			log.Debug(fmt.Sprintf("Token is: %v", token))
			ctx.JSON(http.StatusOK, &tokenResponse{Token: token})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	}
}
