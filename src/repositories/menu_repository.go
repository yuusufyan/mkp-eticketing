package repositories

import (
	"auth-rbac/src/dto"
	"auth-rbac/src/models"

	"gorm.io/gorm"
)

type MenuRepository interface {
	Create(menu *models.Menu) error
	FindAllPaginate(req *dto.PaginateRequest) ([]models.Menu, int64, error)
}

type menuRepository struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) MenuRepository {
	return &menuRepository{db: db}
}

func (r *menuRepository) Create(menu *models.Menu) error {
	return r.db.Create(menu).Error
}

func (r *menuRepository) FindAllPaginate(req *dto.PaginateRequest) ([]models.Menu, int64, error) {
	var menus []models.Menu
	var total int64

	query := r.db.Model(&models.Menu{})

	if req.Search != "" {
		query = query.Where("name ILIKE ?", "%"+req.Search+"%")
	}

	if req.Status == "" {
		query = query.Where("status = ?", "Active")
	} else {
		query = query.Where("status = ?", req.Status)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// pagination offset limit
	offset := (req.Page - 1) * req.PerPage
	if offset < 0 {
		offset = 0
	}
	if req.PerPage == 0 {
		req.PerPage = 10
	}

	if err := query.
		Offset(offset).
		Limit(req.PerPage).
		Find(&menus).Error; err != nil {
		return nil, 0, err
	}

	return menus, total, nil

}
