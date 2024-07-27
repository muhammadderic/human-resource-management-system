package handlers

import "github.com/gin-gonic/gin"

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) SignUpHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "sign up",
	})
}

func (h *AuthHandler) LogInHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "log in",
	})
}
