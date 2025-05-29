package middleware

import (
	"github.com/gin-gonic/gin"
)

type MyMiddleware interface {
	ValidateRequiredHeader(key string) gin.HandlerFunc
}

type Middleware struct {
}

func InitMiddleware() MyMiddleware {
	return &Middleware{}
}
