package main

import (
	"github.com/gin-gonic/gin"
	"serviceAuth/backend/adapter"
	"serviceAuth/backend/middleware"
	"serviceAuth/backend/route"
	"serviceAuth/backend/service/startupService"
)

func main() {

	adapter.InitDatabase()
	startupService.MigrateTenantDB()
	startupService.InitTenantSync()
	adapter.InitTenantDatabase()
	startupService.MigrateDB()
	startupService.PermissionSync()

	engine := gin.New()

	engine.Use(middleware.Ginzap())
	engine.Use(middleware.RecoveryWithZap(true))
	engine.Use(middleware.TenantHandler())

	engine = route.RouteHandler(engine)

	engine.Run()
}
