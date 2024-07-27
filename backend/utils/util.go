package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func BindJSON(c *gin.Context, obj interface{}) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return false
	}
	return true
}

func GenerateHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func SendError(
	c *gin.Context,
	status int,
	message string,
	err error,
) {
	res := gin.H{"error": message}
	if err != nil {
		res["cause"] = err.Error()
	}

	c.JSON(status, res)
}
