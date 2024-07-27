package stores

import (
	"github.com/muhammadderic/hrms/configs"
	"github.com/muhammadderic/hrms/models"
)

type EmployeeRepository interface {
	AddEmployee(employee models.UserPayload) (models.UserPayload, error)
	GetAllEmployees() ([]models.UserPayload, error)
	GetEmployeeById(id string) (*models.UserPayload, error)
}

type EmployeeStore struct{}

func NewEmployeeStore() *EmployeeStore {
	return &EmployeeStore{}
}

func (s *EmployeeStore) AddEmployee(user models.UserPayload) (models.UserPayload, error) {
	result := configs.DB.Create(&user)
	return user, result.Error
}

func (s *EmployeeStore) GetAllEmployees() ([]models.UserPayload, error) {
	users := []models.UserPayload{}
	result := configs.DB.Find(&users)
	return users, result.Error
}

func (s *EmployeeStore) GetEmployeeById(id string) (*models.UserPayload, error) {
	user := models.UserPayload{}
	result := configs.DB.First(&user, id)
	return &user, result.Error
}
