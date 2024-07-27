package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadderic/hrms/stores"
)

type EmployeeHandler struct {
	employeeStore stores.EmployeeRepository
}

func NewEmployeeHandler() *EmployeeHandler {
	return &EmployeeHandler{
		employeeStore: stores.NewEmployeeStore(),
	}
}

func (h *EmployeeHandler) AddEmployee(r *gin.Context) {
}

func (h *EmployeeHandler) GetAllEmployees(r *gin.Context) {
}

func (h *EmployeeHandler) GetEmployee(r *gin.Context) {
}

func (h *EmployeeHandler) UpdateEmployee(r *gin.Context) {
}

func (h *EmployeeHandler) DeleteEmployee(r *gin.Context) {
}
