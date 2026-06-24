package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lucas/confirmation-mariage-app/internal/service"
	"github.com/lucas/confirmation-mariage-app/pkg/response"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// POST /auth/login
func (h *AuthHandler) Login(c *gin.Context) {
	var input service.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, "dados inválidos: "+err.Error())
		return
	}

	result, err := h.authService.Login(input)
	if err != nil {
		response.Unauthorized(c)
		return
	}

	response.OK(c, result)
}
