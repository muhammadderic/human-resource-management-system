package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadderic/hrms/handlers"
)

func RegisterEmployeeRoutes(r *gin.RouterGroup) {
	eh := handlers.NewEmployeeHandler()

	employeeGroup := r.Group("/employee")
	{
		employeeGroup.POST("/", eh.AddEmployee)
		employeeGroup.GET("/", eh.GetAllEmployees)
		employeeGroup.GET("/:id", eh.GetEmployee)
		employeeGroup.PUT("/:id", eh.UpdateEmployee)
		employeeGroup.DELETE("/:id", eh.DeleteEmployee)
	}
}
