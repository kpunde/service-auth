package route

import (
	"github.com/gin-gonic/gin"
	"serviceAuth/backend/controller"
)

func LoginRoutes(route *gin.Engine) {
	routerGroup := route.Group("/auth")
	{
		routerGroup.POST("/login", controller.Login)
	}
}
