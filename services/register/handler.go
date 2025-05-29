package register

import "github.com/gin-gonic/gin"

type RegisterHandler interface {
	RegisterUser(g *gin.Context)
}
