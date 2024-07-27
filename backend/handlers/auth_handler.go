package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/muhammadderic/hrms/models"
	"github.com/muhammadderic/hrms/stores"
	"github.com/muhammadderic/hrms/utils"
	"golang.org/x/crypto/bcrypt"
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
	var body models.UserPayload

	if !utils.BindJSON(c, &body) {
		return
	}

	user := h.authStore.FindUserByEmail(body.Email)
	if user.ID == 0 {
		utils.SendError(
			c,
			http.StatusUnauthorized,
			"Invalid email or password",
			nil,
		)
		return
	}

	err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(body.Password),
	)
	if err != nil {
		c.JSON(
			http.StatusUnauthorized,
			gin.H{"error": "Invalid email or password"},
		)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte("SECRET"))
	if err != nil {
		utils.SendError(
			c,
			http.StatusInternalServerError,
			"Failed to sign token",
			err,
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
