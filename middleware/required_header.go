package middleware

import (
	"auth-service/constants"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m *Middleware) ValidateRequiredHeader(key string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headerVal := ctx.GetHeader(key)
		if headerVal == "" || !constants.AllowedSources[headerVal] {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": fmt.Sprintf("Headers '%s' is required and must be a valid source", key),
			})
		}
		ctx.Set(key, headerVal)
		ctx.Next()
	}
}
