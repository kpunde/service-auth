package utility

import (
	"github.com/gin-gonic/gin"
	"serviceAuth/backend/log"
)

func CheckError(err error) bool {
	if err != nil {
		log.Error(err)
		return true
	}
	return false
}

func CheckErrorResponse(ctx *gin.Context, errorCode int, err error) bool {
	if err != nil {
		log.Error(err)
		ctx.JSON(errorCode, gin.H{"error": err.Error()})
		return true
	}

	return false
}
