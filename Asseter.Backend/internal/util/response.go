package util

import "github.com/gin-gonic/gin"

func GenerateResponse(ctx *gin.Context, code int, message string, isError bool, data any) {
	ctx.JSON(code, gin.H{
		"message": message,
		"code":    code,
		"success": !isError,
		"data":    data,
	})
}
