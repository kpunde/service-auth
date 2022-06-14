package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"serviceAuth/backend/adapter"
	"serviceAuth/backend/log"
	"serviceAuth/backend/utility"
	"strconv"
)

var (
	_tmsService = adapter.NewTenantService()
)

type TenantController interface {
	FindAll(ctx *gin.Context)
	Save(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type tenantController struct {
}

func (t tenantController) FindAll(ctx *gin.Context) {
	ctx.JSON(200, _tmsService.FindAll())
}

func (t tenantController) Save(ctx *gin.Context) {
	var _tenantReq adapter.NewTenantRequest
	err := ctx.ShouldBindJSON(&_tenantReq)
	if !utility.CheckErrorResponse(ctx, http.StatusBadRequest, err) {
		log.Debug(_tenantReq)
		tenant, err := _tmsService.Save(&_tenantReq)
		if !utility.CheckErrorResponse(ctx, 400, err) {
			ctx.JSON(201, tenant)
		}
	}
}

func (t tenantController) Update(ctx *gin.Context) {
	var _tenantReq adapter.NewTenantRequest
	err := ctx.ShouldBindJSON(&_tenantReq)
	if !utility.CheckErrorResponse(ctx, http.StatusBadRequest, err) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
		utility.CheckError(err)
		_tenant := adapter.GetTenantSqlEntity(&_tenantReq)
		_tenant.ID = uint(id)
		if !utility.CheckErrorResponse(ctx, http.StatusBadRequest, err) {
			roleRes, err := _tmsService.Update(_tenant)

			if !utility.CheckErrorResponse(ctx, 400, err) {
				ctx.JSON(http.StatusAccepted, roleRes)
			}
		}
	}
}

func (t tenantController) Delete(ctx *gin.Context) {
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

func NewTenantController() TenantController {
	return tenantController{}
}
