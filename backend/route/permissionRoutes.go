package route

import (
	"github.com/gin-gonic/gin"
	"serviceAuth/backend/controller"
)

var _permissionController = controller.NewPermissionController()

func PermissionRoutes(route *gin.Engine) {
	routerGroup := route.Group("auth/permission")
	{
		routerGroup.GET("/all", _permissionController.FindAll)
	}
}
