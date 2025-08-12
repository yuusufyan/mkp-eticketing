package services

import (
	"auth-rbac/src/dto"
	"auth-rbac/src/models"
	"auth-rbac/src/repositories"
)

type MenuService interface {
	CreateMenu(input dto.CreateMenu) (models.Menu, error)
}

type menuService struct {
	repo repositories.MenuRepository
}

func NewMenuService(r repositories.MenuRepository) MenuService {
	return &menuService{repo: r}
}

func (s *menuService) Create(input dto.CreateMenu) (models.Menu, error) {
	menu := models.Menu{
		Name: input.Name,

		// isi field lain sesuai struct models.Menu kamu
	}
}
