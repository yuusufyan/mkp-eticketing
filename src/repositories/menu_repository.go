package repositories

import (
	"auth-rbac/src/models"

	"gorm.io/gorm"
)

type MenuRepository interface {
	Create(menu *models.Menu) error
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
