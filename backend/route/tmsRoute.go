package route

import (
	"github.com/gin-gonic/gin"
	"serviceAuth/backend/controller"
)

var _tmsController = controller.NewTenantController()

func TmsRoutes(route *gin.Engine) {
	routerGroup := route.Group("auth/tenant")
	{
		routerGroup.GET("/all", _tmsController.FindAll)
		routerGroup.POST("/", _tmsController.Save)
		routerGroup.PUT("/:id", _tmsController.Update)
		routerGroup.DELETE("/:id", _tmsController.Delete)
	}
}
