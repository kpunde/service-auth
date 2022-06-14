package route

import (
	"github.com/gin-gonic/gin"
)

func RouteHandler(engine *gin.Engine) *gin.Engine {
	TmsRoutes(engine)
	HealthCheckRoutes(engine)
	PersonRoutes(engine)
	LoginRoutes(engine)
	RoleRoutes(engine)
	PermissionRoutes(engine)

	return engine
}
