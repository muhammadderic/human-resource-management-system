package main

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadderic/hrms/configs"
	"github.com/muhammadderic/hrms/routes"
)

func main() {
	configs.ConnectDB()

	r := gin.Default()

	apiV1 := r.Group("/api/v1")

	routes.RegisterAuthRoutes(apiV1)

	r.Run(":8080")
}
