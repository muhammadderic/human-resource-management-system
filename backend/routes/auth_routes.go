package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadderic/hrms/handlers"
)

func RegisterAuthRoutes(r *gin.RouterGroup) {
	ah := handlers.NewAuthHandler()

	r.POST("/signup", ah.SignUpHandler)
	r.POST("/login", ah.LogInHandler)
}
