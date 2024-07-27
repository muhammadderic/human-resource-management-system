package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammadderic/hrms/models"
	"github.com/muhammadderic/hrms/stores"
	"github.com/muhammadderic/hrms/utils"
)

type AuthHandler struct {
	authStore stores.AuthRepository
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		authStore: stores.NewAuthStore(),
	}
}

func (h *AuthHandler) SignUpHandler(c *gin.Context) {
	var user models.User

	if !utils.BindJSON(c, &user) {
		return
	}

	hash, err := utils.GenerateHash(user.Password)
	if err != nil {
		utils.SendError(
			c,
			http.StatusInternalServerError,
			"Failed to hash password",
			err,
		)
		return
	}

	user = models.User{
		Email:    user.Email,
		Password: hash,
	}

	result := h.authStore.CreateNewUser(&user)
	if result.Error != nil {
		utils.SendError(
			c,
			http.StatusBadRequest,
			"Failed to create user",
			result.Error,
		)
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *AuthHandler) LogInHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "log in",
	})
}
