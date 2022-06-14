package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"serviceAuth/backend/entity/rest"
	"serviceAuth/backend/entity/sql"
	"serviceAuth/backend/log"
	"serviceAuth/backend/service"
	"serviceAuth/backend/utility"
	"strconv"
)

var (
	_roleService = service.NewRoleService()
)

type RoleController interface {
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Save(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type roleController struct {
}

func (r roleController) FindAll(ctx *gin.Context) {
	reqContext := utility.RequestContextHandler(ctx)
	ctx.JSON(200, _roleService.FindAll(reqContext))
}

func (r roleController) FindById(ctx *gin.Context) {
	reqContext := utility.RequestContextHandler(ctx)
	id := ctx.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 32)
	if !utility.CheckErrorResponse(ctx, 400, err) {
		ppl, err := _roleService.FindById(reqContext, int(idInt))
		if !utility.CheckErrorResponse(ctx, 400, err) {
			ctx.JSON(200, ppl)
		}
	}
}

func (r roleController) Save(ctx *gin.Context) {
	reqContext := utility.RequestContextHandler(ctx)
	var _roleReq rest.NewRoleRequest
	err := ctx.ShouldBindJSON(&_roleReq)
	if !utility.CheckErrorResponse(ctx, http.StatusBadRequest, err) {
		log.Debug(_roleReq)
		role, err := _roleService.Save(reqContext, &_roleReq)
		if !utility.CheckErrorResponse(ctx, 400, err) {
			ctx.JSON(201, role)
		}
	}
}

func (r roleController) Update(ctx *gin.Context) {
	reqContext := utility.RequestContextHandler(ctx)
	var role sql.Role
	err := ctx.ShouldBindJSON(&role)
	if !utility.CheckErrorResponse(ctx, http.StatusBadRequest, err) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
		utility.CheckError(err)
		role.ID = uint(id)
		if !utility.CheckErrorResponse(ctx, http.StatusBadRequest, err) {
			roleRes, err := _roleService.Update(reqContext, &role)

			if !utility.CheckErrorResponse(ctx, 400, err) {
				ctx.JSON(http.StatusAccepted, roleRes)
			}
		}
	}
}

func (r roleController) Delete(ctx *gin.Context) {
	reqContext := utility.RequestContextHandler(ctx)
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	utility.CheckError(err)
	role, err := _roleService.FindById(reqContext, int(id))
	log.Debug(role)
	if !utility.CheckErrorResponse(ctx, http.StatusBadRequest, err) {
		err := _roleService.Delete(reqContext, role)
		if !utility.CheckErrorResponse(ctx, 400, err) {
			ctx.JSON(http.StatusNoContent, nil)
		}
	}
}

func NewRoleController() RoleController {
	return &roleController{}
}
