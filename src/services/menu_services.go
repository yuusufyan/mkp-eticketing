package services

import (
	"auth-rbac/src/dto"
	"auth-rbac/src/models"
	"auth-rbac/src/repositories"

	"github.com/google/uuid"
)

type MenuService interface {
	CreateMenu(input dto.CreateMenu) (models.Menu, error)
	FindAllPaginate(req *dto.PaginateRequest) ([]models.Menu, int64, error)
}

type menuService struct {
	repo repositories.MenuRepository
}

func NewMenuService(r repositories.MenuRepository) MenuService {
	return &menuService{repo: r}
}

// Method sesuai interface
func (s *menuService) CreateMenu(input dto.CreateMenu) (models.Menu, error) {
	menu := models.Menu{
		Name:      input.Name,
		Status:    input.Status,
		ParentID:  input.ParentID,
		CreatedBy: uuid.New().String(),
		UpdatedBy: uuid.New().String(),
	}

	// Simpan ke repository
	if err := s.repo.Create(&menu); err != nil {
		return models.Menu{}, err
	}

	return menu, nil
}

func (s *menuService) FindAllPaginate(req *dto.PaginateRequest) ([]models.Menu, int64, error) {
	// kalau req.Page kosong → set default 1
	if req.Page <= 0 {
		req.Page = 1
	}
	// kalau req.PerPage kosong → set default 10
	if req.PerPage <= 0 {
		req.PerPage = 10
	}

	// lempar ke repository
	menus, total, err := s.repo.FindAllPaginate(req)
	if err != nil {
		return nil, 0, err
	}

	return menus, total, nil
}
