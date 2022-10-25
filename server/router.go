package server

import (
	"belajar-gin/server/controller"

	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.Engine
	user   *controller.UserHandler
}

func NewRouter(router *gin.Engine, user *controller.UserHandler) *Router {
	return &Router{
		router: router,
		user:   user,
	}
}

func (r *Router) Start(port string) {
	user := r.router.Group("/users")
	user.GET("/", r.user.GetUsers)

	r.router.Run()
}
