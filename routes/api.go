package routes

import (
	"auth-service/middleware"
	"auth-service/services/register"

	"github.com/gin-gonic/gin"
)

type route struct {
	g          *gin.Engine
	middleware middleware.MyMiddleware
}

func NewRoute(g *gin.Engine, middleware middleware.MyMiddleware) *route {
	return &route{
		g:          g,
		middleware: middleware,
	}
}

func (r *route) NewRegisterRoutes(handler register.RegisterHandler) {
	api := r.g.Group("/api")

	api.POST("/v1/register", r.middleware.ValidateRequiredHeader("source"), handler.RegisterUser)
}
