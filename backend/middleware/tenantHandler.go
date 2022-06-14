package middleware

import (
	"github.com/gin-gonic/gin"
	"serviceAuth/backend/entity/common"
)

func TenantHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		tenant := context.GetHeader("Tenant")
		if tenant == "" {
			errorWithAbort(context, "Tenant Header not found")
			return
		}
		reqContext := common.RequestContext{Tenant: tenant}
		context.Set("context", reqContext)
	}
}
