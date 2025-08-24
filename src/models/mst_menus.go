package models

import (
	"auth-rbac/src/utils/enums"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Menu struct {
	ID        uuid.UUID          `gorm:"type:uuid;primaryKey"`
	Name      string             `gorm:"size:20;not null"`
	Status    enums.StatusActive `gorm:"type:varchar(20);not null"`
	ParentID  uuid.UUID          `gorm:"type:uuid;nullable"`
	CreatedAt time.Time          `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;not null"`
	CreatedBy string             `gorm:"size:20;not null"`
	UpdatedAt time.Time          `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;not null"`
	UpdatedBy string             `gorm:"size:20;not null"`
}

func (Menu) TableName() string {
	return "mst_menus"
}

func (m *Menu) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return
}
