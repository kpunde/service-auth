package utility

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"serviceAuth/backend/entity/common"
	"serviceAuth/backend/entity/rest"
)

func RequestContextHandler(ctx *gin.Context) *common.RequestContext {
	reqContext, exist := ctx.Get("context")

	if !exist {
		ctx.JSON(http.StatusBadRequest, &rest.Common{Message: "Tenant not found"})
	}
	reqContextObj := reqContext.(common.RequestContext)
	return &reqContextObj
}
