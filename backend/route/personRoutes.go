package route

import (
	"github.com/gin-gonic/gin"
	"serviceAuth/backend/controller"
)

var _personController = controller.NewPersonController()

func PersonRoutes(route *gin.Engine) {
	routerGroup := route.Group("auth/user")
	{
		routerGroup.GET("/all", _personController.FindAll)
		routerGroup.GET("/:id", _personController.FindById)
		routerGroup.POST("/", _personController.Save)
		routerGroup.POST("/:id/update-password", _personController.UpdatePassword)
		routerGroup.PUT("/:id", _personController.Update)
		routerGroup.DELETE("/:id", _personController.Delete)
	}
}
