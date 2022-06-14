package route

import (
	"github.com/gin-gonic/gin"
	"serviceAuth/backend/controller"
)

var _roleController = controller.NewRoleController()

func RoleRoutes(route *gin.Engine) {
	routerGroup := route.Group("auth/role")
	{
		routerGroup.GET("/all", _roleController.FindAll)
		routerGroup.GET("/:id", _roleController.FindById)
		routerGroup.POST("/", _roleController.Save)
		routerGroup.PUT("/:id", _roleController.Update)
		routerGroup.DELETE("/:id", _roleController.Delete)
	}
}
