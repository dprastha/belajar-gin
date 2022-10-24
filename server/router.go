package server

import "github.com/gin-gonic/gin"

type Router struct {
	router *gin.Engine
}

func NewRouter(router *gin.Engine) *Router {
	return &Router{
		router: router,
	}
}

func (r *Router) Start(port string) {
}
