package controller

import (
	"belajar-gin/server/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (u *UserHandler) GetUsers(ctx *gin.Context) {
	response := u.service.GetUsers()

	WriteJsonResponse(ctx, response)
}
