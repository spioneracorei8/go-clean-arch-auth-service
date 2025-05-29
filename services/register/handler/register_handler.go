package handler

import (
	"auth-service/constants"
	"auth-service/services/register"
	"net/http"

	"github.com/gin-gonic/gin"
)

type registerHandler struct {
	registerUs register.RegisterUsecase
}

func NewRegisterHandlerImpl(registerUs register.RegisterUsecase) register.RegisterHandler {
	return &registerHandler{
		registerUs: registerUs,
	}
}

func (h *registerHandler) RegisterUser(g *gin.Context) {
	var (
		source = g.GetHeader("source")
		params map[string]any
	)
	if err := g.ShouldBindJSON(&params); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"msg":   constants.INVALID_REQUEST_BODY,
		})
		return
	}
	if err := h.registerUs.RegisterUser(g.Request.Context(), params, source); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	g.Status(http.StatusCreated)
}
