package route

import (
	"github.com/gin-gonic/gin"
	"serviceAuth/backend/adapter"
	"serviceAuth/backend/middleware"
)

func HealthCheckRoutes(route *gin.Engine) {
	routerGroup := route.Group("auth/", middleware.AuthorizeJWT())
	{
		routerGroup.GET("health", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"is_rest_service_up": true,
				"is_sql_db_conn_up":  adapter.IsSQLDbOkay(),
			})
		})
	}
}
