package controller

import (
	"github.com/gin-gonic/gin"
	"serviceAuth/backend/service"
	"serviceAuth/backend/utility"
)

var (
	_permissionService = service.NewPermissionService()
)

type PermissionController interface {
	FindAll(ctx *gin.Context)
}

type permissionController struct {
}

func (p permissionController) FindAll(ctx *gin.Context) {
	reqContext := utility.RequestContextHandler(ctx)
	ctx.JSON(200, _permissionService.FindAll(reqContext))
}

func NewPermissionController() PermissionController {
	return &permissionController{}
}
