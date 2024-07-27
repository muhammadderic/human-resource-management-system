package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammadderic/hrms/models"
	"github.com/muhammadderic/hrms/stores"
	"github.com/muhammadderic/hrms/utils"
)

type EmployeeHandler struct {
	employeeStore stores.EmployeeRepository
}

func NewEmployeeHandler() *EmployeeHandler {
	return &EmployeeHandler{
		employeeStore: stores.NewEmployeeStore(),
	}
}

func (h *EmployeeHandler) AddEmployee(c *gin.Context) {
	var user models.UserPayload

	if !utils.BindJSON(c, &user) {
		return
	}

	newUser, err := h.employeeStore.AddEmployee(user)
	if !utils.HandleStoreError(c, err) {
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

func (h *EmployeeHandler) GetAllEmployees(c *gin.Context) {
	users, err := h.employeeStore.GetAllEmployees()
	if !utils.HandleStoreError(c, err) {
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *EmployeeHandler) GetEmployee(c *gin.Context) {
	id := c.Param("id")

	user, err := h.employeeStore.GetEmployeeById(id)
	if !utils.HandleNotFoundError(c, err, "User not found") {
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *EmployeeHandler) UpdateEmployee(c *gin.Context) {
	id := c.Param("id")
	var user models.UserPayload

	if !utils.BindJSON(c, &user) {
		return
	}

	updatedUser, err := h.employeeStore.UpdateEmployee(id, user)
	if !utils.HandleStoreError(c, err) {
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

func (h *EmployeeHandler) DeleteEmployee(c *gin.Context) {
	id := c.Param("id")

	if !utils.HandleStoreError(c, h.employeeStore.DeleteEmployee(id)) {
		return
	}

	c.Status(http.StatusNoContent)
}
