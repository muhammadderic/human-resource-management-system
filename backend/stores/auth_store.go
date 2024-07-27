package stores

import (
	"github.com/muhammadderic/hrms/configs"
	"github.com/muhammadderic/hrms/models"
)

type AuthRepository interface {
	CreateNewUser(user *models.User) *configs.DBResult
}

type AuthStore struct{}

func NewAuthStore() *AuthStore {
	return &AuthStore{}
}

func (s *AuthStore) CreateNewUser(user *models.User) *configs.DBResult {
	result := configs.DB.Create(user)

	return &configs.DBResult{
		Result: result,
		Error:  result.Error,
	}
}
