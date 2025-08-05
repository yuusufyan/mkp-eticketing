package handlers

import (
	"auth-rbac/src/repositories"

	"golang.org/x/crypto/bcrypt"
)

type AuthHandler interface {
	Login(username, password string) (string, error)
}

type authHandler struct {
	userRepo repositories.UserRepository
}

func NewAuthHandler(userRepo repositories.UserRepository) AuthHandler {
	return &authHandler{userRepo}
}

func (h *authHandler) Login(username, password string) (string, error) {
	user, err := h.userRepo.FindByUsername(username)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	return user.Role.Name, nil
}
