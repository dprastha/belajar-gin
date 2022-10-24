package controller

import (
	"belajar-gin/server/response"

	"github.com/gin-gonic/gin"
)

func WriteJsonResponse(ctx *gin.Context, payload *response.Response) {
	ctx.JSON(
		payload.Status,
		payload,
	)
}

func WriteErrorJsonResponse(ctx *gin.Context, payload *response.Response) {
	ctx.AbortWithStatusJSON(
		payload.Status,
		payload,
	)
}
