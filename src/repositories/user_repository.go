package repositories

import (
	"auth-rbac/src/config"
	"auth-rbac/src/models"
)

type UserRepository interface {
	FindByUsername(username string) (*models.User, error)
}

type userRepository struct{}

func (r *userRepository) RegisterAccount()

func (r *userRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := config.DB.
		Preload("Role").
		Where("username = ?", username).
		First(&user).Error

	return &user, err
}
